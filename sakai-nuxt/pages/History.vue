<template>
    <div class="card p-fluid" style="height: 100vh">
        <Toolbar class="mb-4">
            <template #start>
                <Button label="Добавить" icon="pi pi-plus" severity="success" class="mr-2" @click="openDialog" />
                <Button label="Остановить" icon="pi pi-trash" severity="danger" @click="confirmDeleteSelected" :disabled="!selectedProducts || !selectedProducts.length" />
            </template>
            <template #end>
                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)" />
            </template>
        </Toolbar>
        <DataTable v-model:editingRows="editingRows" :value="products" v-model:selection="selectedProducts" editMode="row" dataKey="id" @row-edit-save="onRowEditSave" tableClass="editable-cells-table" tableStyle="min-width: 50rem">
            <Column selectionMode="multiple" style="width: 5%" :exportable="false"></Column>
            <Column field="code" header="Дата создания" style="width: 20%">
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>
            <Column field="name" header="Имя" style="width: 20%">
                <template #editor="{ data, field }">
                    <InputText v-model="data[field]" />
                </template>
            </Column>
            <Column field="inventoryStatus" header="Статус" style="width: 20%">
                <template #editor="{ data, field }">
                    <Dropdown v-model="data[field]" :options="statuses" optionLabel="label" optionValue="value" placeholder="Select a Status">
                        <template #option="slotProps">
                            <Tag :value="slotProps.option.value" :severity="getStatusLabel(slotProps.option.value)" />
                        </template>
                    </Dropdown>
                </template>
                <template #body="slotProps">
                    <Tag :value="slotProps.data.inventoryStatus" :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                </template>
            </Column>
            <Column style="width: 10%; min-width: 8rem" bodyStyle="text-align:center">
                <template #body="{ data, field }"> <Button icon="pi pi-eye" text rounded aria-label="Filter" @click="viewDetailt(data)" /></template
            ></Column>
        </DataTable>

        <Dialog v-model:visible="productDialog" :style="{ width: '450px' }" header="Survey Details" :modal="true" class="p-fluid">
            <div class="field">
                <label for="name">Имя</label>
                <InputText id="name" v-model.trim="product.name" required="true" autofocus :class="{ 'p-invalid': submitted && !product.name }" />
                <!-- <small class="p-error" v-if="submitted && !product.name">Name is required.</small> -->
            </div>
            <div v-for="question in questions">
                <div class="field">
                    <InputText placeholder="Вопрос" id="question" v-model="question.description" />
                    <!-- <small class="p-error" v-if="submitted && !product.name">Name is required.</small> -->
                </div>
            </div>
            <Button label="Добавить вопрос" icon="pi pi-plus" @click="addQuestion" />

            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
                <Button label="Save" icon="pi pi-check" text @click="saveProduct" />
            </template>
        </Dialog>
        <Dialog v-model:visible="statisticDialog" :style="{ width: '450px' }" header="Survey Details" :modal="true" class="p-fluid">
            <p>Survey: {{ selectedProduct.name }}</p>
            <p>Status: {{ selectedProduct.inventoryStatus }}</p>
            <div v-for="question in selectedProduct.questions">
                <p>Вопрос: {{ question.description }}</p>
            </div>
            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="statisticDialog = false" />
            </template>
        </Dialog>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ProductService } from '@/service/ProductService';
definePageMeta({
    layout: false
});
const confirmDeleteSelected = () => {
    console.log('selectedProduct:', selectedProducts.value);
    // products.value.filter((val) => val.name != selectedProducts.value[0].name);
    // selectedProducts.value.forEach(element => {
    //     element.
    // });
    for (var i = 0; i < selectedProducts.value.length; i++) {
        products.value[products.value.findIndex((val) => (val.name = selectedProducts.value[i].name))].inventoryStatus = 'НЕАКТИВНО';
    }
};
const selectedProduct = ref();
const statisticDialog = ref(false);
const openDialog = () => {
    productDialog.value = true;
};
const hideDialog = () => {
    productDialog.value = false;
    questions.value = [{ description: '' }];
};
const product = {
    name: 'Имя опроса'
};
const questions = ref([{ description: '' }]);
const selectedProducts = ref();
const productDialog = ref(false);
const products = ref([]);
const editingRows = ref([]);
const statuses = ref([
    { label: 'In Stock', value: 'АКТИВНО' },
    { label: 'Low Stock', value: 'НЕАКТИВНО' },
    { label: 'Out of Stock', value: 'OUTOFSTOCK' }
]);
const addQuestion = () => {
    questions.value.push({ description: '' });
};
onMounted(() => {
    // ProductService.getProductsMini().then((data) => (products.value = data));
    products.value = [{ code: '19-00', name: 'name', inventoryStatus: 'АКТИВНО', questions: [{ description: 'Idk' }] }];
});

const onRowEditSave = (event) => {
    let { newData, index } = event;

    products.value[index] = newData;
};
const currentDate = new Date();
const saveProduct = () => {
    // Get the current time components
    const day = String(currentDate.getDate()).padStart(2, '0'); // Get day and pad with leading zero if needed
    const month = String(currentDate.getMonth() + 1).padStart(2, '0'); // Get month (0-indexed) and pad with leading zero if needed
    const year = currentDate.getFullYear(); // Get full year

    const formattedDate = `${day}.${month}.${year}`;
    products.value.push({ name: product.name, code: formattedDate, inventoryStatus: 'АКТИВНО' });
    hideDialog();
};
const getStatusLabel = (status) => {
    switch (status) {
        case 'АКТИВНО':
            return 'success';

        case 'НЕАКТИВНО':
            return 'warning';

        case 'OUTOFSTOCK':
            return 'danger';

        default:
            return null;
    }
};
const formatCurrency = (value) => {
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' }).format(value);
};
const viewDetailt = (data) => {
    selectedProduct.value = data;
    // console.log(data);
    statisticDialog.value = true;
};
</script>

<style lang="scss" scoped>
::v-deep(.editable-cells-table td.p-cell-editing) {
    padding-top: 0.6rem;
    padding-bottom: 0.6rem;
}
</style>

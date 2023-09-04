<template>
    <div style="margin-bottom: 2rem; align-items: center; margin-left: 10%">
        <FlexColumn>
            <p style="font-weight: bold; font-size: 3dvh">Мысалдар</p>
            <Carousel class="images" :value="products" :numVisible="3" :numScroll="1" :responsiveOptions="responsiveOptions">
                <template #item="data">
                    <img class="image" src="https://www.askify.kz/_next/static/media/example5.cd9a1696.png" />
                </template>
            </Carousel>
        </FlexColumn>
    </div>
</template>
<script>
export default {
    data() {
        return {
            products: [
                {
                    image: 'https://www.askify.kz/_next/static/media/example5.cd9a1696.png'
                },
                {
                    image: 'https://www.askify.kz/_next/static/media/example5.cd9a1696.png'
                },
                {
                    image: 'https://www.askify.kz/_next/static/media/example5.cd9a1696.png'
                },
                {
                    image: 'https://www.askify.kz/_next/static/media/example5.cd9a1696.png'
                },
                {
                    image: 'https://www.askify.kz/_next/static/media/example5.cd9a1696.png'
                }
            ],
            responsiveOptions: [
                {
                    breakpoint: '1199px',
                    numVisible: 3,
                    numScroll: 3
                },
                {
                    breakpoint: '991px',
                    numVisible: 2,
                    numScroll: 2
                },
                {
                    breakpoint: '767px',
                    numVisible: 1,
                    numScroll: 1
                }
            ]
        };
    },
    mounted() {
        // ProductService.getProductsSmall().then((data) => (this.products = data.slice(0, 9)));
    },
    methods: {
        getSeverity(status) {
            switch (status) {
                case 'INSTOCK':
                    return 'success';

                case 'LOWSTOCK':
                    return 'warning';

                case 'OUTOFSTOCK':
                    return 'danger';

                default:
                    return null;
            }
        }
    }
};
</script>
<style scoped>
.images {
    /* style="height: 40%; max-width: 70%" */
    height: 40%;
    max-width: 70%;
}
.image {
    max-width: 90%;
    border-radius: 5%;
    margin-right: 1rem;
    margin-top: 1rem;
}
@media (max-width: 768px) {
    .images {
        height: 100%;
        max-width: 100%;
        align-items: center;
    }
    .image {
        align-self: center;
        max-width: 100%;
        border-radius: 5%;
        /* margin-right: 10%;
        margin-top: 10%; */
    }
}
</style>

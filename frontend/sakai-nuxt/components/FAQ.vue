<template>
    <div>
        <span style="font-size: 3dvh; font-weight: bold">Жиі қойылатын сұрақтар</span>
        <Tree :value="nodes" style="width: 100%; height: 20%; margin-top: 1%; border: none; font-size: 1dvh; background: none; font-size: large" :expanded-keys="expandedKeys" @node-expand="onNodeExpand"></Tree>
    </div>
</template>
<script>
export default {
    data() {
        return {
            expandedKeys: [],
            nodes: [
                {
                    label: 'Бұл веб-қосымша не үшін қажет?',
                    children: [
                        {
                            label: 'Сіздің ойыңыздағы сұраққа бірден жауап беретін қазақ тіліндегі алғашқы жоба. Кез келген сұраққа жауап табасыз🚀'
                        }
                    ]
                },
                {
                    label: 'Неге кейбір сұрақтарға дұрыс жауап бермейді?',
                    children: [
                        {
                            label: `Иә, кейбір сұрақтарға қате жіберіп, толыққанды жауап бере алмауы мүмкін. Оның бірнеше себептері бар:

                                1. Сұрақтың дұрыс қойылмауы;
                                2. Сұрақтың грамматикалық қателермен жазылуы;
                                3. Ашық емес немесе көпке белгісіз ақпаратты сұрау;

                                Сұрақтарды мысалдар бөлімінде көрсетілгендей және сұраныс бетіндегі нұсқаулық бойынша қойыңыз😉`
                        }
                    ]
                }
            ]
        };
    },
    methods: {
        onNodeExpand(e) {
            const key = e.node.key;
            if (this.expandedKeys.includes(key)) {
                // Node is already expanded, so collapse it
                this.expandedKeys = this.expandedKeys.filter((k) => k !== key);
            } else {
                // Node is not expanded, so expand it
                this.expandedKeys.push(key);
            }
        }
    }
};
</script>
<style></style>

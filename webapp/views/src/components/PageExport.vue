<template>
    <b-button v-on:click="doExport" class="float-right icon-only green-tosca">
        <i class="fa fa-fw fa-file-excel"></i>
    </b-button>
</template>

<script>
export default {
    name: "pageExport",
    props: ["storeNames", "rowSelectInvolved"],
    data() {
        return {};
    },
    computed: {
        store () { return this.$store.state.exportData },
        tableStores () { 
            var stores = []; 

            this.storeNames.forEach(name => {
                stores = stores.concat(this.$store.state[name].all);
            });

            return stores;
        },
    },
    methods: {
        doExport(){
            var params = [];
            this.tableStores.forEach((tableStore, i) => {
                tableStore.left.isLoading = true;
                
                var param = tableStore.param;
                param.Filename = tableStore.filename;
                param.Queryname = tableStore.queryname;
                param.Pagination = _.cloneDeep(param.Pagination)
                param.Headers = tableStore.leftHeaders;

                if(!param.Pagination)
                    param.Pagination = {}

                param.Pagination.rowsPerPage = -1;

                param.RowSelect = [];
                if(this.rowSelectInvolved == true){
                    if(tableStore.selected){
                        if(tableStore.selected.length > 0){
                            param.RowSelect = this.$store.getters[this.storeNames[i] + "/getSelectedData"]
                        }
                    }
                }

                params = params.concat(param);
            });

            return this.$store.dispatch(`exportData/doExport`, params).then(() => {
                this.tableStores.forEach((tableStore) => {
                    tableStore.left.isLoading = false;
                })

                this.store.filenames.forEach(filename => {
                    var url = "/csv/" + filename;

                    const a = document.createElement('a');
                    a.style.display = 'none';
                    a.href = url;
                    // the filename you want
                    a.download = filename + ".csv";
                    document.body.appendChild(a);
                    a.click();
                    window.URL.revokeObjectURL(url);
                });
            });
        }
    }
};
</script>
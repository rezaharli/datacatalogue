<style>
    .export-form .column-list{
        max-height: calc(100vh - 117px);
        overflow-y: auto;
    }
</style>

<template>
    <div>
        <b-button @click.stop="drawer = !drawer" class="float-right icon-only green-tosca"><i class="fa fa-fw fa-file-excel"></i></b-button>
        
        <v-navigation-drawer right absolute temporary class="export-form" v-model="drawer">
            <v-card>
                <v-list subheader class="column-list">
                    <v-subheader>Choose fields to export</v-subheader>

                    <v-list-tile v-bind:key="i" v-for="(header, i) in exportableHeaders">
                        <v-list-tile-action>
                            <v-switch selected color="purple" v-model="selectedHeaders[i]" :value="header"></v-switch>
                        </v-list-tile-action>

                        <v-list-tile-title>{{ header.text }}</v-list-tile-title>
                    </v-list-tile>
                </v-list>

                <v-card-actions>
                    <v-spacer></v-spacer>

                    <v-btn color="primary" flat @click="doExport">Export</v-btn>
                </v-card-actions>
            </v-card>
        </v-navigation-drawer>
    </div>
</template>

<script>
export default {
    name: "pageExport",
    props: ["storeName", "rowSelectInvolved"],
    data() {
        var exportableHeaders = [
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Country', value: 'COUNTRY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Itam ID', value: 'ITAM' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Source System Name', value: 'EDM_SOURCE_SYSTEM_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Database Name', value: 'DATABASE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Tier', value: 'TIER' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Certified (Yes/No)', value: 'CERTIFIED' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Name', value: 'TABLE_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Table Description', value: 'TABLE_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Name', value: 'COLUMN_NAME' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Description', value: 'COLUMN_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Type', value: 'DATA_TYPE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Column Length', value: 'COLUMN_LENGTH' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Nullable (Yes/No)', value: 'NULLABLE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Primary Key', value: 'PRIMARY_KEY' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'PII', value: 'PII' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Data Lineage', value: 'DATA_LINEAGE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'CDE (Yes/No)', value: 'CDE' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term', value: 'BUSINESS_TERM' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Term Description', value: 'BUSINESS_DESCRIPTION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Determines Client Location (Yes/No)', value: 'DETERMINES_CLIENT_LOCATION' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Determines Account / Deal Location (Yes/No)', value: 'DETERMINES_ACCOUNT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Business Segment', value: 'BUSINESS_SEGMENT' },
            { align: 'left', display: true, exportable: true, displayCount: false, sortable: true, filterable: true, text: 'Product Category', value: 'PRODUCT_CATEGORY' },
        ];

        return {
            menu: false,
            drawer: false,
            exportableHeaders: _.cloneDeep(exportableHeaders),
            selectedHeaders: _.cloneDeep(exportableHeaders),
        };
    },
    computed: {
        store () { return this.$store.state.exportData },
        tableStore () { 
            return this.$store.state[this.storeName].all;
        },
        headers() {
            return this.tableStore.leftHeaders.filter(v => v.exportable);
        }
    },
    methods: {
        doExport(){
            this.tableStore.left.isLoading = true;
            
            var param = this.tableStore.param;
            param.Filename = this.tableStore.filename;
            param.Queryname = this.tableStore.queryname;
            param.Pagination = this._.cloneDeep(param.Pagination)
            param.Headers = this.selectedHeaders.filter(v => v);

            if( ! param.Pagination)
                param.Pagination = {}

            param.Pagination.rowsPerPage = -1;

            param.RowSelect = [];
            if(this.rowSelectInvolved == true){
                if(this.tableStore.selected){
                    if(this.tableStore.selected.length > 0){
                        param.RowSelect = this.$store.getters[this.storeName + "/getSelectedData"]
                    }
                }
            }

            return this.$store.dispatch(`exportData/doExport`, param).then(() => {
                this.tableStore.left.isLoading = false;

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
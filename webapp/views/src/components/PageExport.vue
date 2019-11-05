<template>
    <!-- <download-excel
        :fields     = "excelFields"
        :fetch      = "fetchData"
        worksheet   = "My Worksheet"
        name        = "filename.xls">
    
        <b-button class="float-right icon-only green-tosca">
            <i class="fa fa-fw fa-file-excel"></i>
        </b-button>
    </download-excel> -->

    <b-button v-on:click="doExport" class="float-right icon-only green-tosca">
        <i class="fa fa-fw fa-file-excel"></i>
    </b-button>
</template>

<script>
import XLSX from 'xlsx'
export default {
    name: "pageExport",
    props: ["storeName", "leftTableCols", "rightTableCols", "forceRightAtFirstLevel", "rowSelectInvolved"],
    data() {
        return {};
    },
    computed: {
        store () { return this.$store.state.exportData },
        tableStore () { return this.$store.state[this.storeName].all },
        excelFields (){
            var ret = {}
            
            _.each(this.LeftTableCols, function(v){
                ret[v.text] = v.value.split(".").reverse()[0];
            })

            if(this.tableStore.isRightTable){
                _.each(this.RightTableCols, function(v){
                    ret[v.text] = v.value.split(".").reverse()[0];
                })
            }

            return ret
        },
        LeftTableCols () {
            var tofilter = [];
            
            if (typeof(this.leftTableCols) == "object") {
                Object.keys(this.leftTableCols).forEach(key => {
                    tofilter = tofilter.concat(this.leftTableCols[key]);
                });
            } else {
                tofilter = this.leftTableCols;
            }
            
            return tofilter.filter(v => v.exportable)
        },
        RightTableCols () {
            var tofilter = [];
            
            if (typeof(this.RightTableCols) == "object") {
                Object.keys(this.leftTableCols).forEach(key => {
                    tofilter = tofilter.concat(this.RightTableCols[key]);
                });
            } else {
                tofilter = this.RightTableCols;
            }
            
            return tofilter.filter(v => v.filterable)
        }
    },
    methods: {
        doExport(){
            this.tableStore.left.isLoading = true;

            var param = this.tableStore.param;
            param.Filename = this.tableStore.filename;
            param.Queryname = this.tableStore.queryname;
            param.Pagination = _.cloneDeep(param.Pagination)
            param.Headers = this.tableStore.leftHeaders;
            
            param.Pagination.rowsPerPage = -1;

            return this.$store.dispatch(`exportData/doExport`, param).then(res => {
                this.tableStore.left.isLoading = false;
                window.open("/csv/"+this.store.filename, "_blank");
            });
        },
        getLeftTable () {
            return this.$store.dispatch(`${this.storeName}/exportData`)
        },
        async fetchData(){
            var fetchExportDatas = async () => {
                if(this.rowSelectInvolved == true){
                    if(this.tableStore.selected){
                        if(this.tableStore.selected.length > 0){
                            return this.tableStore.left.source.filter(v => this.tableStore.selected.find(w => v.TABLE_NAME == w.TABLE_NAME) != undefined); // SEMENTARA (SHOULDN'T BE HARDCODED USING TABLENAME)
                        }
                    }
                }

                const resource = await this.getLeftTable();
                return this.tableStore.exportDatas;
            }

            var exportDatas = await fetchExportDatas();

            var res = [];
            this._.each(exportDatas, (leftRow, i) => {
                var temp = {}

                this.LeftTableCols.forEach(v => {
                    var level = v.value.split(".").length;

                    if(level == 1) {
                        temp[v.value] = leftRow[v.value];
                    } else {
                        var splitted = v.value.split(".");
                        var parentName = splitted[0];
                        var colName = splitted.reverse()[0];

                        // temp[colName] = _.uniq(_.map(leftRow[parentName], colName).filter(Boolean)).join(', ')
                        temp[colName] = leftRow[colName];
                    }
                });
                
                if(this.tableStore.isRightTable){
                    var rightRows = this._.filter(this.tableStore.right.source, (v) => v.LEFTID == leftRow.ID);
                    
                    if(rightRows.length > 0){
                        this._.each(rightRows, (row, i) => {
                            var firstLevel = _.cloneDeep(temp);

                            var firstLevelCols = this.forceRightAtFirstLevel ? this.RightTableCols : _.filter(this.RightTableCols, (v) => v.value.split(".").length == 1);

                            firstLevelCols.forEach(v => {
                                var level = v.value.split(".").length;

                                if(level == 1) {
                                    firstLevel[v.value] = row[v.value];
                                } else {
                                    var splitted = v.value.split(".");
                                    var parentName = splitted[0];
                                    var colName = splitted.reverse()[0];

                                    if(this.forceRightAtFirstLevel) {
                                        firstLevel[colName] = _.uniq(_.map(row[parentName], colName).filter(Boolean)).join(', ');
                                    } else {
                                        firstLevel[colName] = row[colName];
                                    }
                                }
                            });

                            if( ! this.forceRightAtFirstLevel) {
                                var secondLevelCols = _.filter(this.RightTableCols, (v) => v.value.split(".").length == 2);
                                if(secondLevelCols.length > 0){
                                    var colName = secondLevelCols[0].value.split(".")[0];

                                    if(row[colName]){
                                        this._.each(row[colName], (row2, j) => {
                                            var secondLevel = _.cloneDeep(firstLevel);

                                            secondLevelCols.forEach(v => {
                                                secondLevel[v.value.split(".").reverse()[0]] = row2[v.value.split(".").reverse()[0]];
                                            });
                                            
                                            res.push(_.cloneDeep(secondLevel));
                                        })
                                    } else {
                                        res.push(_.cloneDeep(firstLevel));
                                    }
                                } else {
                                    res.push(_.cloneDeep(firstLevel));
                                }
                            } else {
                                res.push(_.cloneDeep(firstLevel));
                            }
                        });
                    } else {
                        res.push(_.cloneDeep(temp));
                    }
                } else {
                    res.push(_.cloneDeep(temp));
                }
            });

            return res
        },
        onExport () {
            this.fetchData().then(res => {
                
                // --- swap object keys with object values ---
                function swap(json){
                    var ret = {};
                    for(var key in json){
                        ret[json[key]] = key;
                    }
                    return ret;
                }
                var excelFieldsSwapped = swap(this.excelFields);

                // --- rename keys with display text ---
                const renameKeys = (keysMap, obj) =>
                    Object.keys(obj).reduce(
                        (acc, key) => ({
                        ...acc,
                        ...{ [keysMap[key] || key]: obj[key] }
                        }),
                        {}
                    );
                var excelDataFinal = [];
                _.each(res, function(v, i){
                    excelDataFinal[i] = renameKeys(excelFieldsSwapped, v);
                });
                
                var excelWS = XLSX.utils.json_to_sheet(excelDataFinal);
                var wb = XLSX.utils.book_new() // make Workbook of Excel
                XLSX.utils.book_append_sheet(wb, excelWS, 'My Worksheet') // sheetAName is name of Worksheet
                XLSX.writeFile(wb, this.storeName+'.xlsx') // name of the file is 'book.xlsx'
            });
        }
    }
};
</script>
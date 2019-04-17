<template>
    <download-excel
        :data     = "excelData"
        :fields   = "excelFields"
        worksheet = "My Worksheet"
        name      = "filename.xls">
    
        <b-button class="float-right icon-only green-tosca">
            <!-- <v-icon dark>file_copy</v-icon> -->
            <i class="fa fa-fw fa-file-excel"></i>
        </b-button>
    </download-excel>
</template>

<script>
export default {
    name: "pageExport",
    props: ["storeName", "leftTableCols", "rightTableCols", "forceRightAtFirstLevel"],
    data() {
        return {};
    },
    computed: {
        store () { return this.$store.state[this.storeName].all },
        excelFields (){
            var ret = {}

            _.each(this.leftTableCols, function(v){
                ret[v.text] = v.value.split(".").reverse()[0];
            })

            if(this.store.isRightTable){
                _.each(this.rightTableCols, function(v){
                    ret[v.text] = v.value.split(".").reverse()[0];
                })
            }

            return ret
        },
        excelData () {
            var res = [];

            this._.each(this.store.left.display, (leftRow, i) => {
                var temp = {}

                this.leftTableCols.forEach(v => {
                    var level = v.value.split(".").length;

                    if(level == 1) {
                        temp[v.value] = leftRow[v.value];
                    } else {
                        var splitted = v.value.split(".");
                        var parentName = splitted[0];
                        var colName = splitted.reverse()[0];

                        temp[colName] = _.uniq(_.map(leftRow[parentName], colName).filter(Boolean)).join(', ')
                    }
                });
                
                if(this.store.isRightTable){
                    var rightRows = this._.filter(this.store.right.display, (v) => v.LEFTID == leftRow.ID);
                    
                    if(rightRows.length > 0){
                        this._.each(rightRows, (row, i) => {
                            var firstLevel = _.cloneDeep(temp);

                            var firstLevelCols = this.forceRightAtFirstLevel ? this.rightTableCols : _.filter(this.rightTableCols, (v) => v.value.split(".").length == 1);

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
                                var secondLevelCols = _.filter(this.rightTableCols, (v) => v.value.split(".").length == 2);
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
        }
    }
};
</script>
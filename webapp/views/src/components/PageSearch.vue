<template>
    <div class="input-group mb-3">
        <input v-model="store.searchMain" type="text" class="form-control" placeholder="Search" aria-label="Recipient's username" aria-describedby="basic-addon2">

        <div class="input-group-append">
            <b-dropdown right id="ddown1" text="" ref="ddownSearch">
                <b-container>
                    <b-form-row class="main-table-search-dropdown-form">
                        <b-col>
                            <b-form @submit="onSubmit" @reset="onReset">
                                <b-form-group v-for="(input, i) in inputs" v-bind:key="i" horizontal :label-cols="4" breakpoint="md" :label="input.label" :label-for="input.source">
                                    <b-form-input type="text" v-if="input.type == 'text'" :id="input.source" v-model="store.searchDropdown[input.source]"></b-form-input>

                                    <b-form-select v-if="input.type == 'dropdown'" :id="input.source" :options="input.options" v-model="store.searchDropdown[input.source]"></b-form-select>
                                </b-form-group>

                                <b-button-group class="mx-1 float-right">
                                    <b-button type="reset" variant="danger">Reset</b-button>
                                    <b-button type="submit" variant="primary">Submit</b-button>
                                </b-button-group>
                            </b-form>
                        </b-col>
                    </b-form-row>
                </b-container>
            </b-dropdown>
        </div>
    </div>
</template>

<script>
export default {
    name: "pageSearch",
    props: ["storeName", "searchDDInputs"],
    data() {
        return {};
    },
    computed: {
        store () { return this.$store.state[this.storeName].all },
        inputs: { 
            get: function(){ return this.searchDDInputs },
            set: function(newVal){
                this.$emit('update:searchDDInputs', newVal)
            }  
        },
        addressPath (){
            var tmp = this.$route.path.split("/")
            return tmp.slice(0, 3).join("/")
        },
    },
    methods: {
        getLeftTable () {
            return this.$store.dispatch(`${this.storeName}/getLeftTable`)
        },
        getRightTable (id) {
            this.$store.dispatch(`${this.storeName}/getRightTable`, id)
        },
        onSubmit (evt) {
            if(evt) evt.preventDefault();

            this.getLeftTable().then(res => {
                var currentLeftID = this.$route.params.system;

                var isLeftTableOwnCurrentRight = this.store.left.display.find(v => v.ID.toString() == currentLeftID);

                if(this.store.isRightTable){
                    if(isLeftTableOwnCurrentRight) { 
                        this.store.isRightTable = true;
                        this.getRightTable(currentLeftID);
                    } else {
                        this.store.isRightTable = false;
                    }
                }

                if( ! this.store.isRightTable){
                    if(isLeftTableOwnCurrentRight) { 
                        this.store.isRightTable = true;
                        this.getRightTable(currentLeftID);
                    } else {
                        this.store.isRightTable = false;
                    }
                }

                if( ! this.store.isRightTable){
                    this.inputs = this.inputs.map((input, i) => {
                        if(input.type == "dropdown") input.options = []; 

                        return input;
                    });
                }
                
                if(this.store.left.display.length == 1){
                    var leftID = this.store.left.display[0].ID;

                    this.$router.push(this.addressPath + "/" + leftID)
                }
            });

            this.$refs.ddownSearch.hide(true);
        },
        onReset (evt) {
            evt.preventDefault();
            /* Reset our form values */
            this.store.searchMain = '';
            this.searchDDInputs.forEach(v => {
                this.store.searchDropdown[v.source] = '';
            });

            this.$router.push(this.addressPath);
            this.onSubmit();

            // /* Trick to reset/clear native browser form validation state */
            // this.searchForm.show = false;
            // this.$nextTick(() => { this.searchForm.show = true });
        },
    }
};
</script>
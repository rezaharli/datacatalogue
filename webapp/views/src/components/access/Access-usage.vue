<style>
  @import '../../assets/styles/dashboard.css';
</style>

<style>
table.v-table thead th > div.btn-group {
    width: auto;
}
</style>

<template>
    <b-row>
        <b-col>
        <!-- Users details -->
            <router-view/>

            <!-- Main content -->
            <b-row>
                <b-col>
                    <v-data-table
                        :headers="headers"
                        :items="roles"
                        class="elevation-1">
                        <template slot="items" slot-scope="props">
                            <td>{{ props.item.RoleName }}</td>
                            <td>HOME</td>
                            <td>Enabled</td>
                        </template>
                    </v-data-table>

                    <!-- <v-data-table
                            :headers="headers"
                            :items="users.items"
                            :loading="users.loading"
                            class="elevation-1">
                        <v-progress-linear slot="progress" color="blue" indeterminate></v-progress-linear>

                        <template slot="no-data">
                            <v-alert :value="!users.loading" color="error" icon="warning">
                                Sorry, nothing to display here :(
                            </v-alert>
                        </template>

                        <template slot="headers" slot-scope="props">
                            <tr>
                                <th v-for="header in props.headers" :key="header.text">
                                    {{ header.text }} ({{ distinctData(header.value, users.items).length }})

                                    <b-dropdown v-if="header.filter"  no-caret variant="link" class="header-filter-icon">
                                        <template slot="button-content"><i class="fa fa-filter text-muted"></i></template>

                                        <b-dropdown-header><b-form-input type="text" placeholder="Filter"></b-form-input></b-dropdown-header>

                                        <b-dropdown-item v-for="item in distinctData(header.value, users.items)" v-bind:key="item">{{ item }}</b-dropdown-item>
                                    </b-dropdown>
                                </th>
                            </tr>
                        </template>

                        <template slot="items" slot-scope="props">
                            <tr :active="props.selected" @click="props.selected = !props.selected">
                                <td>{{ props.item.Username }}</td>
                                <td>{{ props.item.Email }}</td>
                                <td>{{ props.item.Name }}</td>
                                <td>{{ props.item.Role }}</td>
                                <td>{{ props.item.Status }}</td>
                                <td>{{ props.item.CreatedAt }}</td>
                                <td>{{ props.item.UpdatedAt }}</td>
                                <td class="justify-center layout px-0">
                                    <v-icon
                                        small
                                        class="mr-2"
                                        @click="editItem(props.item)"
                                    >edit</v-icon>

                                    <v-icon
                                        small
                                        @click="deleteItem(props.item)"
                                    >delete</v-icon>
                                </td>
                            </tr>
                        </template>
                    </v-data-table> -->
                </b-col>
            </b-row>
        </b-col>
    </b-row>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
    data () {
        return {
            headers: [
                { text: 'Role Name', align: 'left', value: 'RoleName', sortable: false, filter: true },
                { text: 'Landing Page', align: 'left', value: 'RoleName', sortable: false, filter: true },
                { text: 'Status', align: 'left', value: 'RoleName', sortable: false, filter: true },
            ],
            roles: [
                { RoleName: 'Admin' },
                { RoleName: 'DSC' },
                { RoleName: 'DPO' },
                { RoleName: 'DDO' },
                { RoleName: 'RFO' },
            ]
        }
    }
}
</script>

<script setup lang="ts">
import { computed, onMounted, ref } from 'vue';
import Pagination from '@/components/PaginationData.vue';
import Filters from '@/components/FilterTable.vue';
import ExportData from '@/components/ExportData.vue';
import ImportData from './ImportData.vue';
import { usePeopleData } from '@/services/peopleData';
import { useAuthStore } from '@/stores/auth'

const { data, totalRecords, page, itemsPerPage, loading, fetchData, selectedPage } = usePeopleData();
const authStore = useAuthStore()

const headers = ref([
  { title: 'ID', key: 'id', sortable: true,},
  { title: 'Age', key: 'age', sortable: true },
  { title: 'Education', key: 'education', sortable: true },
  { title: 'Marital Status', key: 'marital_status', sortable: true },
  { title: 'Occupation', key: 'occupation', sortable: true },
  { title: 'Income', key: 'income', sortable: true }
]);

const pageCount = computed(() => Math.ceil(totalRecords.value / itemsPerPage.value))

const handlePageChange = (newPage: number) => {
  if (page.value !== newPage) {
    page.value = newPage;
    selectedPage.value = newPage;
    fetchData();
  } else {
    fetchData(true);
  }
};

const handleItemsPerPageChange = (newItemsPerPage: number) => {
  if (itemsPerPage.value !== newItemsPerPage) {
    itemsPerPage.value = newItemsPerPage;
    fetchData();
  }
};

// Fetch initial data
onMounted(() => {
  fetchData();
});
</script>

<template>
  <p class="text-blue-400 mb-4 text-xl flex-1 flex justify-start items-start">Welcome {{ authStore.username }}!</p>
  <p>This page is designed to manage data from a dataset created by Barry Becker based on a 1994 census containing over 30,000 records with demographic data, essential for understanding what types of people earned more or less than $50,000 USD, among other insights.</p>
  <Filters @update:page="handlePageChange"/>
  <div class="flex flex-col justify-end space-x-4 mb-4">
    <ImportData />
    <ExportData />
  </div>
  <v-data-table
    :headers="headers"
    :items="data"
    :loading="loading"
    v-model:page="page"
    :items-per-page="itemsPerPage"
  >    
    <template v-slot:bottom>
      <Pagination 
      :page="page" 
      :page-count="pageCount" 
      :items-per-page="itemsPerPage" 
      @update:page="handlePageChange"
      @update:items-per-page="handleItemsPerPageChange"
      />
    </template>
    <template v-slot:[`item.id`]="{ item }">
      {{ item.id }}
    </template>
    <template v-slot:[`item.age`]="{ item }">
      {{ item.age }}
    </template>
    <template v-slot:[`item.education`]="{ item }">
      {{ item.education }}
    </template>
    <template v-slot:[`item.marital_status`]="{ item }">
      {{ item.marital_status }}
    </template>
    <template v-slot:[`item.occupation`]="{ item }">
      {{ item.occupation }}
    </template>
    <template v-slot:[`item.income`]="{ item }">
      {{ item.income }}
    </template>
  </v-data-table>
</template>

<style scoped>
.v-data-table thead th {
  background-color: #007bff; /* Color azul */
  color: white; /* Color del texto */
}
</style>
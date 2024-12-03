// src/services/peopleData.ts
import { apiHandleData, apiDownload, apiUploadFile } from '@/services/utils';
import { ref } from 'vue';
import { useFiltersStore } from '@/stores/filters';
import type { PeopleData } from '@/types/index'; 

export function usePeopleData() {
  const filtersStore = useFiltersStore();

  const data = ref<PeopleData[]>([]);
  const totalRecords = ref(0);
  const page = ref(1);
  const selectedPage = ref(1);
  const itemsPerPage = ref(10);
  const loading = ref(false);

  async function fetchData (resetPage = false, exportData = false, graphicData = false) {
    loading.value = true;

    if (resetPage) {
      selectedPage.value = 1;
    }

    const params : Record<string, string> = {
      page: exportData ? "": graphicData ? "" : String(selectedPage.value),
      limit: exportData ? "" : graphicData ? "" : String(itemsPerPage.value),
      export: exportData ? 'true' : 'false',
    };
    const filterState = filtersStore.$state;

    Object.entries(filterState).forEach(([key, value]) => {
      if (key === 'min_age' || key === 'max_age') {
        // Solo incluir si tiene un valor válido
        if (value !== null && value !== undefined) {
          params[key] = String(value);
        }
      } else if (value !== null && value !== undefined && value !== '') {
        params[key] = String(value);
      }
    });

    try {
      const query = new URLSearchParams(params).toString();

      if (exportData) {
        const blob = await apiDownload(`/people-data?${query}`)
        const url = window.URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'people_data.csv');
        document.body.appendChild(link);
        link.click();
        link.remove();
        return;
      } else{
        const response = await apiHandleData(`/people-data?${query}`);
        data.value = response.data;
        totalRecords.value = response.total;
      }

    } catch (error) {
      console.error('Failed to fetch data', error);
    } finally {
      loading.value = false;
    }
  };

  async function importData(file: File) {
    try {
      await apiUploadFile(`/people-data/updateData`, file);
      fetchData(true);
    } catch (error) {
      console.error('Failed to upload file', error);
    }
  }

  return { data, totalRecords, page, itemsPerPage, loading, fetchData, selectedPage, importData };
}

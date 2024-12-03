import { apiHandleData } from '@/services/utils';
import { useFiltersStore } from '@/stores/filters';

export function useFilters() {
  const filtersStore = useFiltersStore();

  async function saveFilters() {
    try {
      const params: Record<string, string> = {};
      const state = filtersStore.$state;

      Object.entries(state).forEach(([key, value]) => {
        if (key === 'min_age' || key === 'max_age') {
          // Solo incluir si tiene un valor válido
          if (value && value !== null && value !== undefined) {
            params[key] = String(value);
          }
        } else if (value !== null && value !== undefined && value !== '') {
          params[key] = String(value);
        }
      });

      const query = new URLSearchParams(params).toString();

      await apiHandleData(`/filters/save?${query}`, {
        method: 'POST',
      });
    } catch (error) {
      console.error('Failed to save filters:', error);
    }
  }

  async function loadFilters() {
    try {
      const response = await apiHandleData('/filters/load');

      if (response) {
        filtersStore.setFilters(response.data);
        return response.data;
      } else{
        filtersStore.setFilters({});
        return null;
      }
      

    } catch (error) {
      console.error('Failed to load filters:', error);
      return null;
    }
  }

  async function loadFilterOptions() {
    try {
        const response = await apiHandleData('/filters/options');
        return response;
    } catch (error) {
        console.error('Failed to load filter options:', error);
        return null;
    }
}

  return { saveFilters, loadFilters, loadFilterOptions };
}
import { defineStore } from 'pinia';
import { ref } from 'vue';

export const useFiltersStore = defineStore('filters', () => {
  const education = ref<string | null | undefined>(null);
  const marital_status = ref<string | null | undefined>(null);
  const occupation = ref<string | null | undefined>(null);
  const min_age = ref<number | null>(null);
  const max_age = ref<number | null>(null);
  const income = ref<string | null | undefined>(null);

  const setFilters = (newFilters: any) => {
    if (!newFilters) {
      resetFilters();
      return;
    }
    education.value = newFilters.education ?? null;
    marital_status.value = newFilters.marital_status || null;
    occupation.value = newFilters.occupation || null;
    min_age.value = newFilters?.min_age !== undefined ? Number(newFilters.min_age) : 0;
    max_age.value = newFilters?.max_age !== undefined ? Number(newFilters.max_age) : 100;
    income.value = newFilters.income || null;
  };

  const resetFilters = () => {
    education.value = null;
    marital_status.value = null;
    occupation.value = null;
    min_age.value = 0;
    max_age.value = 100;
    income.value = null;
  };

  return {
    education,
    marital_status,
    occupation,
    min_age,
    max_age,
    income,
    setFilters,
    resetFilters
  };
});

import { defineStore } from 'pinia';
import { ref } from 'vue';

export const usePaginationStore = defineStore('pagination', () => {
  // Estado reactivo
  const selectedPage = ref(1);
  const currentPage = ref(1);

  // Acción para sincronizar la página actual y seleccionada
  const setCurrentPage = (page: number) => {
    currentPage.value = page;
    selectedPage.value = page;
  };

  // Retorna las propiedades y métodos del store
  return {
    selectedPage,
    currentPage,
    setCurrentPage,
  };
});
<!-- trucode_data_management_app/src/components/BasicChart.vue -->
<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { usePeopleData } from '@/services/peopleData';
import { UserCharts } from '@/services/charts';
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import { TitleComponent, TooltipComponent, LegendComponent } from "echarts/components";
import VChart from "vue-echarts";

use([ CanvasRenderer, PieChart, TitleComponent, TooltipComponent, LegendComponent ]);


// Obtener datos de ingresos y actualizar el gráfico
const { data, fetchData, loading } = usePeopleData();
const { createBarChartOptions, updateBarChartOptions, resizeChart } = UserCharts();

const option = createBarChartOptions("Marital status distribution"); 

const chartInstance = ref<any>(null);

onMounted(async () => {
  await fetchData(false, false, true); // Llama a la función para obtener los datos
  updateChartIncomeData(); // Actualiza los datos del gráfico

  window.addEventListener('resize', resizeChart);
});

onBeforeUnmount(() => {
  // Limpiar el evento de resize cuando el componente se destruya
  window.removeEventListener('resize', resizeChart);
});

// Función para actualizar los datos del gráfico
const updateChartIncomeData = () => {
  updateBarChartOptions(data.value, option.value); // Usar la función del servicio para actualizar los datos
};

</script>

<template>
  <v-chart 
  v-if="!loading && data.length > 0"
  ref="chartInstance"
  class="chart w-full h-96" 
  :option="option" />
  <div v-if="loading && data.length > 0" class="flex flex-col items-center w-full">
    <div class="mt-4 text-center text-gray-600">Loading...</div>
    <v-progress-circular indeterminate color="primary" />
  </div>
</template>

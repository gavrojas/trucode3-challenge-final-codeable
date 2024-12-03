<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue';
import { usePeopleData } from '@/services/peopleData';
import { UserCharts } from '@/services/charts';
import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import { BarChart } from 'echarts/charts';
import { TitleComponent, TooltipComponent, LegendComponent, GridComponent } from "echarts/components";
import VChart from "vue-echarts";

use([ CanvasRenderer, PieChart, BarChart, TitleComponent, TooltipComponent, LegendComponent, GridComponent ]);

const { createStackedBarChartOptions, updateStackedBarChartData, resizeChart } = UserCharts();

const option = createStackedBarChartOptions("Education by marital status distribution"); // Usar la función del servicio

const { data, fetchData, loading } = usePeopleData();
const chartInstance = ref<any>(null);

onMounted(async () => {
  await fetchData(false, false, true); // Llama a la función para obtener los datos
  updateChartData(); // Actualiza los datos del gráfico

  window.addEventListener('resize', () => resizeChart(chartInstance));
});

onBeforeUnmount(() => {
  // Limpiar el evento de resize cuando el componente se destruya
  window.removeEventListener('resize', () => resizeChart(chartInstance));
});

// Función para actualizar los datos del gráfico
const updateChartData = () => {
  updateStackedBarChartData(data.value, option.value); // Usar la función del servicio para actualizar los datos
};
</script>

<template>
  <v-chart 
    v-if="!loading && data.length > 0"
    ref="chartInstance"
    class="chart w-full h-96" 
    :option="option" />
  <div v-if="loading" class="flex flex-col items-center w-full">
    <div class="mt-4 text-center text-gray-600">Loading...</div>
    <v-progress-circular indeterminate color="primary" />
  </div>
</template>
<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { usePeopleData } from '@/services/peopleData';
import VChart from 'vue-echarts';
import { BarChart } from 'echarts/charts';
import { TitleComponent, TooltipComponent, LegendComponent, GridComponent } from 'echarts/components';
import { use } from 'echarts/core';

use([BarChart, TitleComponent, TooltipComponent, LegendComponent, GridComponent]);

const averageAge = ref(0);
const totalGreaterThan50K = ref(0);
const totalLessThan50K = ref(0);
const incomeByEducationChartOptions = ref({});
const incomeByOccupationChartOptions = ref({});
const { data, fetchData } = usePeopleData(); 

// Definir tipos para los conteos de ingresos
type IncomeDistribution = {
  '<=50K': number;
  '>50K': number;
};

const calculateStatistics = () => {
  if (data.value.length === 0) return;

  // Calcular la edad promedio
  const totalAge = data.value.reduce((sum, person) => sum + (person.age || 0), 0);
  averageAge.value = totalAge / data.value.length;

  totalGreaterThan50K.value = 0;
  totalLessThan50K.value = 0;


  const incomeByEducation: Record<string, IncomeDistribution> = {};
  const incomeByOccupation: Record<string, IncomeDistribution> = {};

  data.value.forEach(person => {
    const education = person.education;
    const occupation = person.occupation;
    const income = person.income;

    // Contar ingresos
    if (income === '<=50K') {
      totalLessThan50K.value += 1;
    } else if (income === '>50K') {
      totalGreaterThan50K.value += 1;
    }

    if (education) {
      if (!incomeByEducation[education]) {
        incomeByEducation[education] = { '<=50K': 0, '>50K': 0 };
      }
      if (income === '<=50K') {
        incomeByEducation[education]['<=50K'] += 1;
      } else if (income === '>50K') {
        incomeByEducation[education]['>50K'] += 1;
      }
    }

    // Contar por ocupación
    if (occupation) {
      if (!incomeByOccupation[occupation]) {
        incomeByOccupation[occupation] = { '<=50K': 0, '>50K': 0 };
      }
      if (income === '<=50K') {
        incomeByOccupation[occupation]['<=50K'] += 1;
      } else if (income === '>50K') {
        incomeByOccupation[occupation]['>50K'] += 1;
      }
    }
  });

  // Configurar opciones de gráficos
  incomeByEducationChartOptions.value = createIncomeDistributionChartOptions(incomeByEducation, '');
  incomeByOccupationChartOptions.value = createIncomeDistributionChartOptions(incomeByOccupation, '');
};

const createIncomeDistributionChartOptions = (data: Record<string, IncomeDistribution>, title: string) => {
  const categories = Object.keys(data);
  const lessThan50KCounts = categories.map(category => data[category]['<=50K']);
  const greaterThan50KCounts = categories.map(category => data[category]['>50K']);

  return {
    title: {
      text: title,
      left: 'center'
    },
    tooltip: {
      trigger: 'axis'
    },
    legend: {
      data: ['<=50K', '>50K'],
      bottom: 0
    },
    xAxis: {
      type: 'category',
      data: categories
    },
    yAxis: {
      type: 'value'
    },
    series: [
      {
        name: '<=50K',
        type: 'bar',
        data: lessThan50KCounts
      },
      {
        name: '>50K',
        type: 'bar',
        data: greaterThan50KCounts
      }
    ]
  };
};

onMounted(async () => {
  await fetchData(false, false, true);; 
  calculateStatistics(); // Calcular estadísticas después de obtener los datos
});
</script>

<template>
  <div class="mb-4">
    <h2 class="text-lg font-semibold mb-4">Demographic Statistics</h2>
    <p>The following data and charts represent the data you can find in the 'view data' table. If you want to change any filter, you can do so from that view. <br> For better visualization of the charts, view them on a tablet or computer screen.</p>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <div class="p-4 bg-blue-100 rounded-lg shadow">
        <h3 class="text-md font-semibold">Income Distribution by Education</h3>
        <v-chart v-if="incomeByEducationChartOptions" :option="incomeByEducationChartOptions" class="chart w-full h-48" />
        <p class="mt-2 text-sm text-gray-600">
          This chart shows the distribution of income levels among different educational backgrounds. 
          It helps to understand how education impacts income potential.
        </p>
      </div>
      <div class="p-4 bg-green-100 rounded-lg shadow">
        <h3 class="text-md font-semibold">Income Distribution by Occupation</h3>
        <v-chart v-if="incomeByOccupationChartOptions" :option="incomeByOccupationChartOptions" class="chart w-full h-48" />
        <p class="mt-2 text-sm text-gray-600">
          This chart illustrates the income distribution across various occupations. 
          It provides insights into how different job roles affect income levels.
        </p>
      </div>
      <div class="p-4 bg-purple-100 rounded-lg shadow">
        <h3 class="text-md font-semibold">Average age:</h3>
        <p class="text-xl">{{ averageAge.toFixed(3) }}</p>
        <h3 class="text-md font-semibold">Total people earning greater than 50K</h3>
        <p class="text-xl">{{ totalGreaterThan50K }}</p>
        <h3 class="text-md font-semibold">Total People earning less than 50k</h3>
        <p class="text-xl">{{ totalLessThan50K }}</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.p-4 bg-gray-100 rounded-lg shadow {
  padding: 16px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
.chart {
  width: 100%;
  height: 200px;
}
</style>
// services/charts.ts
import { ref } from 'vue';
import type { ChartOption } from '@/types/index';
import type { PeopleData } from '@/types/index';

export function UserCharts() {
  // PieChart
  const createPieChartOptions = () => {
    const option = ref<ChartOption>({
      title: {
        text: "Income distribution",
        left: "center"
      },
      tooltip: {
        trigger: "item",
        formatter: "{a} <br/>{b} : {c} ({d}%)"
      },
      legend: {
        orient: "horizontal",
        bottom: 'bottom',
        data: [] // Inicialmente vacío, se llenará con los datos
      },
      series: [
        {
          name: "Income",
          type: "pie",
          radius: "55%",
          center: ["50%", "60%"],
          data: [], // Inicialmente vacío, se llenará con los datos
          emphasis: {
            itemStyle: {
              shadowBlur: 10,
              shadowOffsetX: 0,
              shadowColor: "rgba(0, 0, 0, 0.5)"
            }
          }
        }
      ]
    });
  
    return option;
  };
  
  const updatePieChartOptions = (data: PeopleData[], option: ChartOption) => {
    const incomeData = [
      { value: 0, name: "<=50K" },
      { value: 0, name: ">50K" }
    ];
  
    // Contar la cantidad de cada ingreso
    data.forEach(PeopleData => {
      if (PeopleData.income === "<=50K") {
        incomeData[0].value += 1;
      } else if (PeopleData.income === ">50K") {
        incomeData[1].value += 1;
      }
    });
  
    // Actualizar los datos en la opción del gráfico
    option.series[0].data = incomeData;
    option.legend.data = incomeData.map(item => item.name); 
  };
  // Bar chart
  const createBarChartOptions = (title: string) => {
    const option = ref<ChartOption>({
      title: {
        text: title,
        left: "center"
      },
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "shadow"
        }
      },
      legend: {
        orient: 'horizontal',
        data: [], // Inicialmente vacío, se llenará con los datos
        bottom: 0,
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '16%',
        containLabel: true
      },
      xAxis: {
        type: "category",
        data: [] // Inicialmente vacío, se llenará con los datos
      },
      yAxis: {
        type: "value"
      },
      series: [] // Inicialmente vacío, se llenará con los datos
    });
    return option;
  }

  const updateBarChartOptions = (data: PeopleData[], option: ChartOption) => {
    if (!option.xAxis || !option.legend || !option.series) {
      throw new Error("Option is not properly defined");
    }

    const maritalStatusCount: Record<string, number> = {};
    // Contar el estado civil
    data.forEach(person => {
      if (person.marital_status) {
        maritalStatusCount[person.marital_status] = (maritalStatusCount[person.marital_status] || 0) + 1;
      }
    });
    const maritalStatuses = [...new Set(data.map(PeopleData => PeopleData.marital_status))];
    // Actualizar la opción del gráfico
    option.xAxis.data = Object.keys(maritalStatusCount); // Asignar estados civiles al eje X
    option.series = [{
      name: 'Count',
      type: 'bar',
      data: Object.values(maritalStatusCount),
      }]
    option.legend.data = maritalStatuses; // Asignar estados civiles a la leyenda
  }

  // StackedBarChart
  const createStackedBarChartOptions = (title: string) => {
    const option = ref<ChartOption>({
      title: {
        text: title,
        left: "center"
      },
      tooltip: {
        trigger: "axis",
        axisPointer: {
          type: "shadow"
        }
      },
      legend: {
        orient: 'horizontal',
        data: [], // Inicialmente vacío, se llenará con los datos
        bottom: 0,
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '16%',
        containLabel: true
      },
      xAxis: {
        type: "category",
        data: [] // Inicialmente vacío, se llenará con los datos
      },
      yAxis: {
        type: "value"
      },
      series: [] // Inicialmente vacío, se llenará con los datos
    });
  
    return option;
  };

  const updateStackedBarChartData = (data: PeopleData[], option: ChartOption) => {
    if (!option.xAxis || !option.legend) {
      throw new Error("Option is not properly defined");
    }
  
    const educationLevels = [...new Set(data.map(PeopleData => PeopleData.education))]; // Obtener niveles educativos únicos
    const maritalStatuses = [...new Set(data.map(PeopleData => PeopleData.marital_status))]; // Obtener estados civiles únicos
  
    // Inicializar la serie de datos
    const seriesData = maritalStatuses.map(status => {
      return {
        name: status,
        type: "bar",
        stack: "total",
        data: educationLevels.map(level => {
          // Contar cuántas personas tienen este nivel educativo y estado civil
          return data.filter(PeopleData => PeopleData.education === level && PeopleData.marital_status === status).length;
        })
      };
    });
  
    // Actualizar la opción del gráfico
    option.xAxis.data = educationLevels; // Asignar niveles educativos al eje X
    option.legend.data = maritalStatuses; // Asignar estados civiles a la leyenda
    option.series = seriesData; // Asignar datos de la serie
  };
  
  const resizeChart = (chartInstance: any) => {
    if (chartInstance.value) {
      // Redimensiona el gráfico de ECharts
      chartInstance.value.resize();
    }
  };

  const createLineChartOptions = (title: string) => {
    const option = ref<ChartOption>({
      title: {
        text: title,
        left: "center"
      },
      tooltip: {
        trigger: "axis"
      },
      legend: {
        data: [],
        bottom: 0,
      },
      grid: {
        left: '3%',
        right: '4%',
        bottom: '16%',
        containLabel: true
      },
      xAxis: {
        type: "category",
        data: [] // Inicialmente vacío, se llenará con los datos
      },
      yAxis: {
        type: "value"
      },
      series: [] // Inicialmente vacío, se llenará con los datos
    });
    
    return option;
  }

  const updateLineChartOptions = (data: PeopleData[], option: ChartOption) => {
    if (!option.xAxis || !option.legend || !option.series) {
      throw new Error("Option is not properly defined");
    }
     // Agrupar ingresos por edad
    const incomeByAge: Record<number, { lessThan50K: number; greaterThan50K: number }> = {};
    data.forEach(person => {
      const age = person.age; 
      const income = person.income; 
        if (age !== undefined && income !== undefined) {
          if (!incomeByAge[age]) {
            incomeByAge[age] = { lessThan50K: 0, greaterThan50K: 0 };
          }
           // Contar según la categoría de ingreso
          if (income === "<=50K") {
            incomeByAge[age].lessThan50K += 1;
          } else if (income === ">50K") {
            incomeByAge[age].greaterThan50K += 1;
          }
        }
    });
     // Calcular el ingreso promedio por edad
    const ages = Object.keys(incomeByAge).map(Number);
    const lessThan50KCounts = ages.map(age => incomeByAge[age].lessThan50K);
    const greaterThan50KCounts = ages.map(age => incomeByAge[age].greaterThan50K);
    // Asignar datos al gráfico
    option.xAxis.data = ages; // Asignar edades al eje X
    option.series = [
      {
        name: '<=50K',
        type: 'line',
        data: lessThan50KCounts, // Asignar conteo de <=50K
      },
      {
        name: '>50K',
        type: 'line',
        data: greaterThan50KCounts, // Asignar conteo de >50K
      },
    ];
     // Actualizar la leyenda
    option.legend.data = ['<=50K', '>50K']; // Asignar nombre a la leyenda
  }

  return { createPieChartOptions, updatePieChartOptions, createBarChartOptions, updateBarChartOptions, createStackedBarChartOptions, updateStackedBarChartData, createLineChartOptions, updateLineChartOptions, resizeChart };
}


  


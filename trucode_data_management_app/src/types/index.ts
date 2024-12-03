export interface JWTPayload {
  MapClaims: {
    eat: number
    iat: number
  }
  session: string
}

export interface OptionsApiCall {
  method?: string
  data?: any
  headers?: any
  notifyLogout?: boolean
}

export interface PeopleData {
  id: number
  age: number
  education: string
  marital_status: string
  occupation: string
  income: string
}

interface PieData {
  value: number;
  name: string;
}

export interface ChartOption {
  title?: { text: string; left?: string };
  tooltip?: { trigger: string; formatter?: string, axisPointer?: { type: string }};
  legend: { orient?: string; data: string[]; right?: string | number ; left?: string | number ; bottom?: string | number; top?: string | number; };
  grid?: { left?: string; right?: string; bottom?: string; containLabel?: boolean;} ;
  xAxis?: { type: string | number; data?: string[] | number[];} ;
  yAxis?: { type: string; } ;
  series:  Array<{
    name: string;
    type: string;
    stack?: string; // Solo para gráficos de barras apiladas
    radius?: string; // Solo para gráficos de pastel
    center?: string[]; // Solo para gráficos de pastel
    data: Array<number | PieData>;
    emphasis?: { itemStyle: { 
      shadowBlur?: number; 
      shadowOffsetX?: number; 
      shadowColor?: string 
    }};
  }>;
}
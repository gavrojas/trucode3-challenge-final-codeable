<script setup lang="ts">
import IncomePie from '@/components/charts/IncomePie.vue';
import StackBarEducation from '@/components/charts/EducationMSStackBar.vue';
import MaritalStatusBar from '@/components/charts/MaritalStatusBar.vue';
import LineIncome from '@/components/charts/AgeLineChart.vue';
import DemographicStatistics from '@/components/DemographicStatistics.vue';
</script>

<template>
  <v-container class="flex flex-col gap-12">
    <DemographicStatistics />
    <IncomePie />
    <MaritalStatusBar />
    <StackBarEducation />
    <LineIncome />
  </v-container>
</template>
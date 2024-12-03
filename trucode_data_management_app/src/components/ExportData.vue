<script setup lang="ts">
import { usePeopleData } from '@/services/peopleData';

const { fetchData } = usePeopleData();

const exportData = () => {
  fetchData(false, true, false);
}
</script>

<template>
  <div class="d-flex justify-end">
    <v-btn color="warning" @click="exportData">Export data</v-btn>
  </div>
</template>

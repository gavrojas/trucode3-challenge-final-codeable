<script setup lang="ts">
import { useFiltersStore } from '@/stores/filters';
import { useFilters } from '@/services/filters';
import { onMounted, ref, computed } from 'vue';

const { saveFilters, loadFilters, loadFilterOptions } = useFilters();
const filtersStore = useFiltersStore();

const ageRange = computed({
    get() {
      const minAge = filtersStore.min_age ? Number(filtersStore.min_age) : null;
      const maxAge = filtersStore.max_age ? Number(filtersStore.max_age) : null;

      if (minAge === null && maxAge === null) {
          return [];
      }

      return [minAge ?? 0, maxAge ?? 100];
    },
    set(newRange: number[]) {
      if (Array.isArray(newRange) && newRange.length === 2) {
        filtersStore.min_age = newRange[0];
        filtersStore.max_age = Math.max(newRange[0], newRange[1]);
      }
      return newRange;
    }
});

const filterOptions = ref({
    education: [] as string[],
    marital_status: [] as string[],
    occupation: [] as string[],
    income: [] as string[]
});

const emit = defineEmits(['update:page']);

const applyFilters = async () => {
  await saveFilters();
  emit('update:page', 1);
};

const resetFilters = async () => {
  filtersStore.resetFilters();
  await saveFilters();
  emit('update:page', 1);
};

onMounted(async () => {
  try {
    await loadFilters();

    const options = await loadFilterOptions();
    if (options) {
        filterOptions.value = options
    }
  } catch (error) {
    console.error('Error loading saved filters:', error);
  }
});

</script>

<template>
  <v-card class="mb-4">
    <v-card-title>Filters</v-card-title>
    <v-card-text>
      <v-row>
        <v-col cols="12" md="4">
          <v-select
            v-model="filtersStore.education"
            :items="filterOptions.education"
            label="Education"
            variant="underlined"
            clearable
          />
        </v-col>
        <v-col cols="12" md="4">
          <v-select
            v-model="filtersStore.marital_status"
            :items="filterOptions.marital_status"
            label="Marital Status"
            variant="underlined"
            clearable
          />
        </v-col>
        <v-col cols="12" md="4">
          <v-select
            v-model="filtersStore.occupation"
            :items="filterOptions.occupation"
            label="Occupation"
            variant="underlined"
            clearable
          />
        </v-col>
        <v-col cols="12" md="4">
          <v-select
            v-model="filtersStore.income"
            :items="filterOptions.income"
            label="Income"
            variant="underlined"
            clearable
          />
        </v-col>
        <v-col cols="12">
          <div class="w-full">
            <div class="text-subtitle-1 mb-4">Age Range</div>
            <div class="flex flex-col md:flex-row items-center space-y-4 md:space-y-0 md:space-x-4">
              <v-text-field
                v-model="filtersStore.min_age"
                type="number"
                variant="underlined"
                density="comfortable"
                label="Min Age"
                :max="Number(filtersStore.max_age) || 100"
                :min="0"
                class="w-full md:w-20"
                hide-details
              />
              <v-range-slider
                v-model="ageRange"
                :min="0"
                :max="100"
                :step="1"
                :strict="true"
                thumb-label="always"
                class="w-full px-3"
                hide-details
              />
              <v-text-field
                v-model="filtersStore.max_age"
                type="number"
                variant="underlined"
                density="comfortable"
                label="Max Age"
                :min="Number(filtersStore.min_age) || 0"
                :max="100"
                class="w-full md:w-20"
                hide-details
              />
            </div>
          </div>
        </v-col>
        <v-col cols="12" class="d-flex justify-end">
          <v-btn @click="applyFilters" color="primary">Apply Filters</v-btn>
          <v-btn @click="resetFilters" color="secondary" class="ml-2">Reset</v-btn>
        </v-col>
      </v-row>
    </v-card-text>
  </v-card>
</template>
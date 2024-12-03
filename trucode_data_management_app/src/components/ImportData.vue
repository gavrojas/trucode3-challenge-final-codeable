<script setup lang="ts">
import { ref } from 'vue'
import { usePeopleData } from '@/services/peopleData'

const { importData } = usePeopleData()
const file = ref<File | null>(null)
const showSuccessSnackbar = ref(false);
const successMessage = ref('');
const snackbarTimeout = ref(5000);

const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    file.value = target.files[0]
    showSuccessSnackbar.value = false;
  }
}

const handleImport = async () => {
  if (file.value) {
    try {
      await importData(file.value)
      successMessage.value = 'Data imported and updated successfully!';
      showSuccessSnackbar.value = true;
    } catch (error) {
      console.error('Failed to import data:', error);
    }
  } else {
    alert('Please select a CSV file.');
  }
}
</script>

<template>
  <div class="d-flex flex-col items-end mb-4">
    <v-file-input 
      label="Import your .csv"
      variant="underlined"
      @change="handleFileChange"
      accept=".csv"
      class="w-full max-w-[280px] mb-2"
      counter
      show-size
      >
    </v-file-input>
    <v-btn color="success" @click="handleImport" :disabled="!file" class="max-w-min">Update or Import new data</v-btn>
  </div>
  <v-snackbar
      v-model="showSuccessSnackbar"
      :timeout="snackbarTimeout"
      color="success"
      border="left"
      rounded
      class="flex items-end justify-end"
    > 
      {{ successMessage }}
      <template #actions>
        <v-btn
          color="grey-darken-3"
          variant="text"
          @click="showSuccessSnackbar = false"
        >
          Close
        </v-btn>
      </template>
    </v-snackbar>
</template>
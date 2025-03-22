<script setup>

import SelectModels from "@/components/SelectModels.vue";
import {defineEmits, onBeforeMount, ref} from "vue";

let options = ref([])
onBeforeMount(() => {
  console.log(process.env.VUE_APP_BACKEND_URL + '/models')
  fetch(process.env.VUE_APP_BACKEND_URL + '/models', {
    method: "GET",
    mode: "cors",
     headers: {
      'Content-Type': 'application/json',
     },
  }).then(response => response.json()).then(responseData => {
    options.value = responseData.data
  })
})

const emit = defineEmits(['select'])

const emitSelect = (value) => {
  emit("select", value)
}
</script>

<template>
<SelectModels :options="options" :key="options" @select="emitSelect"/>
</template>

<style scoped>
</style>
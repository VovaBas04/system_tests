<script setup>

import SelectModels from "@/components/SelectModels.vue";
import {onBeforeMount} from "vue";

let options = []
onBeforeMount(() => {
  console.log(process.env.VUE_APP_BACKEND_URL + '/models')
  fetch(process.env.VUE_APP_BACKEND_URL + '/models', {
    method: "GET",
    mode: "no-cors",
    headers: {
      'Content-Type': 'application/json', // Тип содержимого
    },
  }).then(response => response.json()).then(data => {
    console.log(data)
    options = data
  })
})
</script>

<template>
<SelectModels :options="options"/>
</template>

<style scoped>

</style>
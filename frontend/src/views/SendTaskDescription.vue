<script setup>
import TaskTextArea from "@/components/TaskTextArea.vue";
import SendButton from "@/components/SendButton.vue";
import TestsSample from "@/components/TestsSample.vue";
import {defineEmits, reactive, ref} from "vue";
import SelectView from "@/views/SelectView.vue";

const isHidden = reactive({
  value : true
})

const emit = defineEmits(["send", "update:tests"]);
let model = null
let text = ''
let error = ref('')
let fields = [
  { field1: '', field2: '' },
  { field1: '', field2: '' },
  { field1: '', field2: '' },
  { field1: '', field2: '' },
  { field1: '', field2: '' }
]

const sendTask = () => {
  isHidden.value = true
  if (!model) {
    error.value = 'Выберите модель'
    return
  }

  if (!text) {
    error.value = 'Введите текст'

    return
  }
  error.value = ''
  emit("send")
  fetch(process.env.VUE_APP_BACKEND_URL + '/tasks', {
    method : "POST",
    mode: "cors",
    body: JSON.stringify({
      model : model,
      prompt: text,
    }),
    headers: {
      'Content-Type': 'application/json',
    },
  }).then((response) => {
    if (response.status !== 200) {
      error.value = "Ошибка сервера. Переформулируйте запрос и попробуйте еще."
      
      return
    }
    return response.json()
  }).then(data => {
    if (data) {
    fields.value = []
    for (let el of data) {
      fields.value.push({
        "field1": Object.entries(el).filter(([key]) => key !== "answer").map(([key, value]) => `${key}=${value}`).join(';'),
        "field2": el["answer"]
      })
    }
    return true
  }
  }).then((bool) => {
    if (bool) {
      emit('update:tests', fields.value)
      isHidden.value = false
    }
  })
}

const setModel = (data) => {
  model = data.name
}

const setText = (data) => {
  text = data
}

const getTests = (data) => {
  emit("update:tests", data)
}
</script>

<template>
  <SelectView @select = "setModel"/>
  <TaskTextArea
                @update:model-value="setText"
                label="Задача"
                placeholder="Введите текст задачи..."
                :maxLength="1000"
                class="taskText"
                :error="error"/>
  <SendButton @click="sendTask"/>
  <TestsSample :hidden="isHidden.value ?? true" :ref="isHidden" :fields="fields.value" @update:fields="getTests"/>
</template>

<style scoped>
.taskText{
  position: relative;
  top: 100px;
}
</style>
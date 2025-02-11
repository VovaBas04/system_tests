<script setup>
import TaskTextArea from "@/components/TaskTextArea.vue";
import SendButton from "@/components/SendButton.vue";
import TestsSample from "@/components/TestsSample.vue";
import {reactive, defineEmits} from "vue";

const isHidden = reactive({
  value : true
})

const emit = defineEmits(["send"]);

const sendTask = () => {
  emit("send")
  isHidden.value = false
}

</script>

<template>
  <TaskTextArea
                label="Задача"
                placeholder="Введите текст задачи..."
                :maxLength="1000"
                class="taskText"/>
  <SendButton @click="sendTask"/>
  <TestsSample :hidden="isHidden.value ?? true" :ref="isHidden"/>
</template>

<style scoped>
.taskText{
  position: relative;
  top: 100px;
}
</style>
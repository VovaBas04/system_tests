<template>
  <div class="code-editor" :style="{ height: editorHeight + 'px' }">
    <textarea
        v-model="code"
        @input="highlightCode"
        :style="{ height: editorHeight + 'px' }"
    ></textarea>
    <pre class="my-top"><code ref="highlight" class="language-python"></code></pre>
  </div>
</template>

<script setup>
import { ref, onMounted, defineEmits} from 'vue';
import Prism from 'prismjs';
import 'prismjs/themes/prism.css';
import 'prismjs/components/prism-python';

const emit = defineEmits(["update:code"])
const code = ref('');
const highlight = ref(600);
const editorHeight = ref(600);

const highlightCode = () => {
  if (highlight.value) {
    emit("update:code", code.value)
    highlight.value.textContent = code.value;
    Prism.highlightElement(highlight.value);
  }
};

onMounted(() => {
  highlightCode();
});
</script>

<style scoped>
.code-editor {
  pointer-events: all;
  position: relative;
  width: 95%;
  height: 20%;
  font-family: 'Fira Code', monospace;
}

textarea {
  width: 100%;
  height: 100%;
  resize: vertical;
  padding-top: 10px;
  line-height: 1.5;
  border: 1px solid #ccc;
  border-radius: 4px;
  background-color: transparent;
  color: transparent;
  caret-color: #000;
  z-index: 2;
}

pre {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  margin: 0;
  padding: 10px;
  font-size: 14px;
  line-height: 1.5;
  pointer-events: none;
  overflow: hidden;
  background-color: rgba(245, 245, 245, 0) !important;
  font-family: 'Arial', sans-serif;
}

code {
  display: block;
  overflow-x: auto;
  padding: 0;
}

</style>
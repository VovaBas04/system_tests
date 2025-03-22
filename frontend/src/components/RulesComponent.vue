<template>
  <ul class="rules-list">
    <li v-for="(rule, index) in rules" :key="index" class="rule-item">
      <span class="rule-number">{{ index + 1 }})</span>
      {{ rule }}
    </li>
  </ul>
</template>

<script>
export default {
  name: 'RulesComponent',
  props: {
    rules: {
      type: Array,
      required: true
    }
  }
}
</script>

<style scoped>
.rules-list {
  position: absolute;
  top: 10%;
  left: 2%;
  list-style-type: none;
  padding: 0;
}

.rule-item {
  background-color: #e8f4f8;
  margin-bottom: 30px;
  padding: 40px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  transition: transform 0.2s ease-in-out;
}

.rule-item:hover {
  transform: translateY(-3px);
}

.rule-number {
  font-weight: bold;
  color: #3498db;
  margin-right: 10px;
}
</style>


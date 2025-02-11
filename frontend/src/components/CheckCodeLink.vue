<template>
  <div class="link-wrapper">
    <a
        href="#"
        class="animated-link"
        @mousemove="handleMouseMove"
        @mouseleave="handleMouseLeave"
        @click.prevent="handleClick"
    >
      <span class="link-content">
        <span class="link-text">Проверить код</span>
        <span class="link-question">?</span>
      </span>
      <div
          class="glow"
          :style="{
          left: `${glowPosition.x}px`,
          top: `${glowPosition.y}px`
        }"
      ></div>
      <svg class="link-underline" width="100%" height="100%" viewBox="0 0 1200 60" preserveAspectRatio="none">
        <path class="path" d="M0,56.5c0,0,298.666,0,399.333,0C448.336,56.5,513.994,46,597,46c77.327,0,135,10.5,200.999,10.5c95.996,0,402.001,0,402.001,0"></path>
      </svg>
    </a>

    <!-- Декоративные элементы -->
    <div class="particles">
      <span v-for="n in 6" :key="n" class="particle"></span>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const glowPosition = ref({ x: 0, y: 0 })

const handleMouseMove = (event) => {
  const rect = event.currentTarget.getBoundingClientRect()
  glowPosition.value = {
    x: event.clientX - rect.left,
    y: event.clientY - rect.top
  }
}

const handleMouseLeave = () => {
  glowPosition.value = { x: 0, y: 0 }
}

const handleClick = () => {
  // Здесь можно добавить логику при клике
  console.log('Checking code...')
}
</script>

<style scoped>
.link-wrapper {
  display: inline-block;
  position: relative;
  padding: 20px;
  font-family: 'Arial', sans-serif;
}

.animated-link {
  position: relative;
  display: inline-block;
  padding: 10px 20px;
  text-decoration: none;
  font-size: 24px;
  font-weight: 600;
  color: #2c3e50;
  overflow: hidden;
  transition: color 0.3s ease;
}

.link-content {
  position: relative;
  z-index: 2;
  display: flex;
  align-items: center;
  gap: 4px;
}

.link-text {
  background: linear-gradient(to right, #2c3e50, #3498db);
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  transition: all 0.3s ease;
}

.link-question {
  display: inline-block;
  color: #3498db;
  transform: translateY(-2px);
  transition: all 0.3s ease;
}

/* Подчеркивание SVG */
.link-underline {
  position: absolute;
  bottom: 0;
  left: 0;
  pointer-events: none;
  width: 100%;
  height: 10px;
}

.path {
  fill: none;
  stroke: #3498db;
  stroke-width: 1;
  stroke-dasharray: 1200;
  stroke-dashoffset: 1200;
  transition: stroke-dashoffset 0.5s ease;
}

/* Эффект свечения */
.glow {
  position: absolute;
  width: 60px;
  height: 60px;
  background: radial-gradient(circle, rgba(52, 152, 219, 0.3) 0%, rgba(52, 152, 219, 0) 70%);
  border-radius: 50%;
  pointer-events: none;
  transform: translate(-50%, -50%);
  z-index: 1;
  opacity: 0;
  transition: opacity 0.3s ease;
}

/* Частицы */
.particles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
}

.particle {
  position: absolute;
  width: 4px;
  height: 4px;
  background: #3498db;
  border-radius: 50%;
  opacity: 0;
  pointer-events: none;
}

/* Hover эффекты */
.animated-link:hover {
  color: #3498db;
}

.animated-link:hover .link-text {
  background: linear-gradient(to right, #3498db, #2980b9);
  -webkit-background-clip: text;
  background-clip: text;
  transform: translateY(-2px);
}

.animated-link:hover .link-question {
  transform: translateY(-4px) rotate(12deg);
  color: #2980b9;
}

.animated-link:hover .glow {
  opacity: 1;
}

.animated-link:hover .path {
  stroke-dashoffset: 0;
  animation: draw 1s ease forwards;
}

.animated-link:hover .particle {
  animation: particle-animation 0.6s ease-out forwards;
}

/* Анимация частиц */
.animated-link:hover .particle:nth-child(1) {
  top: 20%; left: 20%;
  animation-delay: 0.1s;
}
.animated-link:hover .particle:nth-child(2) {
  top: 60%; left: 80%;
  animation-delay: 0.2s;
}
.animated-link:hover .particle:nth-child(3) {
  top: 40%; left: 40%;
  animation-delay: 0.3s;
}
.animated-link:hover .particle:nth-child(4) {
  top: 80%; left: 60%;
  animation-delay: 0.4s;
}
.animated-link:hover .particle:nth-child(5) {
  top: 30%; left: 70%;
  animation-delay: 0.5s;
}
.animated-link:hover .particle:nth-child(6) {
  top: 70%; left: 30%;
  animation-delay: 0.6s;
}

@keyframes draw {
  0% {
    stroke-dashoffset: 1200;
  }
  100% {
    stroke-dashoffset: 0;
  }
}

@keyframes particle-animation {
  0% {
    transform: translate(0, 0) scale(1);
    opacity: 0;
  }
  50% {
    opacity: 1;
  }
  100% {
    transform: translate(
        calc(var(--direction-x, 1) * 50px),
        calc(var(--direction-y, 1) * 50px)
    ) scale(0);
    opacity: 0;
  }
}

/* Активное состояние */
.animated-link:active {
  transform: scale(0.98);
}

/* Доступность */
@media (prefers-reduced-motion: reduce) {
  .animated-link,
  .link-text,
  .link-question,
  .glow,
  .path,
  .particle {
    transition: none;
    animation: none;
  }
}

/* Фокус для клавиатурной навигации */
.animated-link:focus {
  outline: none;
  box-shadow: 0 0 0 3px rgba(52, 152, 219, 0.3);
  border-radius: 4px;
}

.animated-link:focus:not(:focus-visible) {
  box-shadow: none;
}
</style>
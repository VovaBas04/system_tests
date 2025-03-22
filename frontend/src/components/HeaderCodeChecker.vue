<template>
  <header class="header" :class="{ 'nav-open': isMenuOpen }">
    <div class="header-container">
      <!-- Logo -->
      <div class="logo">
        <span class="logo-text">Code<span class="logo-accent">Checker</span></span>
      </div>

      <!-- Desktop Navigation -->
      <nav class="nav-desktop">
        <a href="/" class="nav-link">Главная</a>
        <a href="/rules" class="nav-link">Правила</a>
        <a href="/docs" class="nav-link">Документация</a>
      </nav>

      <!-- Auth Buttons - Desktop -->
      <div class="auth-buttons">
        <button class="btn btn-login">Log In</button>
        <button class="btn btn-signup">Sign Up</button>
      </div>

      <!-- Mobile Menu Button -->
      <button
          class="mobile-menu-btn"
          @click="toggleMenu"
          :aria-expanded="isMenuOpen"
          aria-label="Toggle menu"
      >
        <span class="hamburger-line"></span>
        <span class="hamburger-line"></span>
        <span class="hamburger-line"></span>
      </button>
    </div>

    <!-- Mobile Navigation -->
    <nav class="nav-mobile">
      <a href="#features" class="nav-link" @click="closeMenu">Features</a>
      <a href="#pricing" class="nav-link" @click="closeMenu">Pricing</a>
      <a href="#docs" class="nav-link" @click="closeMenu">Documentation</a>
      <a href="#about" class="nav-link" @click="closeMenu">About</a>
      <div class="mobile-auth-buttons">
        <button class="btn btn-login">Log In</button>
        <button class="btn btn-signup">Sign Up</button>
      </div>
    </nav>
  </header>
</template>

<script setup>
import { ref } from 'vue'

const isMenuOpen = ref(false)

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value
  if (isMenuOpen.value) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
}

const closeMenu = () => {
  isMenuOpen.value = false
  document.body.style.overflow = ''
}
</script>

<style scoped>
:root {
  --primary-color: #2D2D2D;
  --accent-color: #00FF9B;
  --text-color: #FFFFFF;
  --background-color: #1A1A1A;
  --hover-color: #3A3A3A;
}

.header {
  background-color: var(--primary-color);
  padding: 1rem 0;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
  z-index: 1000;
}

.header-container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 2rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* Logo Styles */
.logo {
  z-index: 100;
}

.logo-text {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--text-color);
}

.logo-accent {
  color: var(--accent-color);
}

/* Navigation Styles */
.nav-desktop {
  display: flex;
  gap: 2rem;
}

.nav-link {
  color: var(--text-color);
  text-decoration: none;
  font-size: 1rem;
  transition: color 0.3s ease;
}

.nav-link:hover {
  color: var(--accent-color);
}

/* Button Styles */
.btn {
  padding: 0.5rem 1rem;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn-login {
  background: transparent;
  color: var(--text-color);
  border: 1px solid var(--text-color);
  margin-right: 1rem;
}

.btn-login:hover {
  border-color: var(--accent-color);
  color: var(--accent-color);
}

.btn-signup {
  background: var(--accent-color);
  color: var(--primary-color);
}

.btn-signup:hover {
  opacity: 0.9;
  transform: translateY(-1px);
}

/* Auth Buttons Container */
.auth-buttons {
  display: flex;
  align-items: center;
}

/* Mobile Menu Button */
.mobile-menu-btn {
  display: none;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  z-index: 100;
}

.hamburger-line {
  display: block;
  width: 24px;
  height: 2px;
  background-color: var(--text-color);
  transition: all 0.3s ease;
  margin: 4px 0;
}

/* Mobile Navigation */
.nav-mobile {
  display: none;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background-color: var(--primary-color);
  padding: 6rem 2rem 2rem;
  transform: translateX(100%);
  transition: transform 0.3s ease;
}

.mobile-auth-buttons {
  margin-top: 2rem;
  display: none;
}

/* Mobile Menu Open States */
.nav-open .nav-mobile {
  transform: translateX(0);
}

.nav-open .hamburger-line:nth-child(1) {
  transform: translateY(6px) rotate(45deg);
}

.nav-open .hamburger-line:nth-child(2) {
  opacity: 0;
}

.nav-open .hamburger-line:nth-child(3) {
  transform: translateY(-6px) rotate(-45deg);
}

/* Media Queries */
@media (max-width: 768px) {
  .nav-desktop,
  .auth-buttons {
    display: none;
  }

  .mobile-menu-btn {
    display: block;
  }

  .nav-mobile {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2rem;
  }

  .mobile-auth-buttons {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    width: 100%;
    max-width: 200px;
  }

  .mobile-auth-buttons .btn {
    width: 100%;
    margin: 0;
  }
}
</style>
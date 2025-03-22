import { createRouter, createWebHistory } from 'vue-router'
import MyHome from "@/views/MyHome.vue";
import RulesPage from "@/views/RulesPage.vue";
const routes = [
  { path: '/', component: MyHome},
  { path: '/rules', component: RulesPage}
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

export default router

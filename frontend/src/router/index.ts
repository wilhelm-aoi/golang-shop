import { createRouter, createWebHistory } from 'vue-router';
import Login from '../components/Login.vue';
import ProductList from '../components/ProductList.vue';
import ProductDetail from '../components/ProductDetail.vue';
import Cart from '../components/Cart.vue';
import UserProfile from '../components/UserProfile.vue';

const routes = [
  { path: '/', component: ProductList },
  { path: '/login', component: Login },
  { path: '/product/:id', component: ProductDetail },
  { path: '/cart', component: Cart },
  { path: '/profile', component: UserProfile },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
import { createStore } from 'vuex';
import axios from 'axios';

export default createStore({
  state: {
    user: null,
    products: [],
    cart: [],
  },
  mutations: {
    setUser(state, user) {
      state.user = user;
    },
    setProducts(state, products) {
      state.products = products;
    },
    addToCart(state, product) {
      state.cart.push(product);
    },
    clearCart(state) {
      state.cart = [];
    },
  },
  actions: {
    async login({ commit }, credentials) {
      const response = await axios.post('/api/login', credentials);
      commit('setUser', response.data);
    },
    async fetchProducts({ commit }) {
      const response = await axios.get('/api/products');
      commit('setProducts', response.data);
    },
    async fetchProduct(_, productId) {
      const response = await axios.get(`/api/products/${productId}`);
      return response.data;
    },
    async fetchUserProfile({ commit }) {
      const response = await axios.get('/api/user/profile');
      commit('setUser', response.data);
    },
    async fetchUserOrders() {
      const response = await axios.get('/api/user/orders');
      return response.data;
    },
    addToCart({ commit }, product) {
      commit('addToCart', product);
    },
    checkout({ commit }) {
      // Simulate checkout process
      commit('clearCart');
    },
  },
});
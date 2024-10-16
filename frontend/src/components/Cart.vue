<template>
  <div>
    <h1>Shopping Cart</h1>
    <el-table :data="cartItems" style="width: 100%">
      <el-table-column prop="name" label="Product Name" width="180"></el-table-column>
      <el-table-column prop="price" label="Price" width="180"></el-table-column>
      <el-table-column prop="quantity" label="Quantity" width="180"></el-table-column>
    </el-table>
    <el-button type="primary" @click="checkout">Checkout</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'vuex';

export default defineComponent({
  name: 'Cart',
  setup() {
    const store = useStore();
    const cartItems = computed(() => store.state.cart);

    const checkout = () => {
      store.dispatch('checkout');
    };

    return {
      cartItems,
      checkout,
    };
  },
});
</script>
<template>
  <div>
    <h1>User Profile</h1>
    <p>Username: {{ user.username }}</p>
    <p>Role: {{ user.role }}</p>
    <h2>Order History</h2>
    <el-table :data="orders" style="width: 100%">
      <el-table-column prop="productName" label="Product Name" width="180"></el-table-column>
      <el-table-column prop="quantity" label="Quantity" width="180"></el-table-column>
      <el-table-column prop="totalPrice" label="Total Price" width="180"></el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { useStore } from 'vuex';

export default defineComponent({
  name: 'UserProfile',
  setup() {
    const user = ref({});
    const orders = ref([]);
    const store = useStore();

    onMounted(async () => {
      user.value = await store.dispatch('fetchUserProfile');
      orders.value = await store.dispatch('fetchUserOrders');
    });

    return {
      user,
      orders,
    };
  },
});
</script>
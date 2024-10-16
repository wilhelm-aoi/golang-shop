<template>
  <div>
    <img :src="product.imageURL" alt="Product Image" />
    <h1>{{ product.name }}</h1>
    <p>{{ product.description }}</p>
    <p>Price: ${{ product.price }}</p>
    <el-button type="primary" @click="addToCart">Add to Cart</el-button>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import { useRoute } from 'vue-router';

export default defineComponent({
  name: 'ProductDetail',
  setup() {
    const product = ref({});
    const store = useStore();
    const route = useRoute();

    onMounted(async () => {
      const productId = route.params.id;
      product.value = await store.dispatch('fetchProduct', productId);
    });

    const addToCart = () => {
      store.dispatch('addToCart', product.value);
    };

    return {
      product,
      addToCart,
    };
  },
});
</script>
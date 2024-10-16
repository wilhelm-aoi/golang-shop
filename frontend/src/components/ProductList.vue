<template>
  <el-row :gutter="20">
    <el-col :span="6" v-for="product in products" :key="product.id">
      <el-card :body-style="{ padding: '0px' }">
        <img :src="product.imageURL" class="image" />
        <div style="padding: 14px;">
          <span>{{ product.name }}</span>
          <div class="bottom clearfix">
            <el-button type="text" @click="viewProduct(product.id)">View</el-button>
          </div>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'ProductList',
  setup() {
    const products = ref([]);
    const store = useStore();
    const router = useRouter();

    onMounted(async () => {
      products.value = await store.dispatch('fetchProducts');
    });

    const viewProduct = (id: number) => {
      router.push(`/product/${id}`);
    };

    return {
      products,
      viewProduct,
    };
  },
});
</script>

<style scoped>
.image {
  width: 100%;
  display: block;
}
</style>
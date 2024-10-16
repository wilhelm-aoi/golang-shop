<template>
  <el-form @submit.prevent="handleLogin">
    <el-form-item>
      <el-input v-model="username" placeholder="Username"></el-input>
    </el-form-item>
    <el-form-item>
      <el-input v-model="password" type="password" placeholder="Password"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" native-type="submit">Login</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';
import { useStore } from 'vuex';
import { useRouter } from 'vue-router';

export default defineComponent({
  name: 'Login',
  setup() {
    const username = ref('');
    const password = ref('');
    const store = useStore();
    const router = useRouter();

    const handleLogin = async () => {
      await store.dispatch('login', { username: username.value, password: password.value });
      router.push('/');
    };

    return {
      username,
      password,
      handleLogin,
    };
  },
});
</script>
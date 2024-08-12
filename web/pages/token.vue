<script setup lang="ts">
import {useTokenState} from "@/store";
import LoadingLayout from "@/layouts/LoadingLayout.vue";

const route = useRoute();

const {code} = route.query;
const tokenState = useTokenState();

onMounted(async () => {
  await tokenState.fetchToken(code as string);
  await navigateTo("/dashboard");
})
</script>

<template>
  <LoadingLayout loading />
</template>

<style scoped lang="scss">
</style>
<script setup lang="ts">
import {useUserState} from "@/store";
import LoadingLayout from "~/layouts/LoadingLayout.vue";

const userState = useUserState();

onMounted(async () => {
  await userState.fetchUser();
});
</script>

<template>
  <LoadingLayout :loading="userState.loading">
    <Error v-if="userState.error" message="We've got something broken" />
    <Error v-else-if="userState.user == null" message="Unfortunately, you can't be here without auth" />
    <slot v-else />
  </LoadingLayout>
</template>

<style scoped lang="scss">
pre {
  color: var(--accent);
}
</style>
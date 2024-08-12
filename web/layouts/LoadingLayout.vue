<script setup lang="ts">
const props = defineProps<{
  loading: boolean
}>()
</script>

<template>
  <div v-if="props.loading" class="loading">
    <img src="/HACKERMANS.webp" alt="HACKERMANS">
    <p>Wait a minute, we're hacking microwaves...</p>
  </div>
  <slot v-else />
</template>

<style scoped lang="scss">
  .loading {
    width: 100%;
    height: 100vh;

    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 12px;

    img {
      width: 48px;
      height: 48px;
    }

    p {
      color: var(--foreground-400);

      font: {
        size: 16px;
        weight: 500;
      };

      line-height: 24px;
    }
  }
</style>
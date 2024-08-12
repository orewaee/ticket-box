<script setup lang="ts">
const props = defineProps<{
  title: string
  description: string
}>()
</script>

<template>
  <div class="option">
    <div class="info">
      <h4>{{title}}</h4>
      <p>{{description}}</p>
    </div>
    <div class="icon">
      <slot />
    </div>
  </div>
</template>

<style scoped lang="scss">
  .option {
    display: flex;
    justify-content: space-between;

    background-color: var(--background-300);
    border-radius: 24px;

    padding: 24px;

    .info {
      font: {
        size: 16px;
        weight: 600;
      };

      line-height: 24px;

      h4 {
        color: var(--foreground-400);
      }

      p {
        color: var(--foreground-100);
      }
    }

    .icon {
      width: 48px;
      height: 48px;

      padding: 12px;
    }
  }
</style>
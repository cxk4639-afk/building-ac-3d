<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="菜单管理">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        菜单权限入口，后续支持目录、菜单、按钮权限树配置。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>

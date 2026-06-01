<script setup lang="ts">
defineProps<{
  embedded?: boolean
}>()
</script>

<template>
  <n-card title="角色管理">
    <n-space vertical size="large">
      <n-alert type="info" title="工程骨架">
        租户权限入口，后续支持角色授权、菜单授权、按钮权限。
      </n-alert>
      <n-data-table :columns="[]" :data="[]" />
    </n-space>
  </n-card>
</template>

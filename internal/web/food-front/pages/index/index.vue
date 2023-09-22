<template>
	<view>
		<u-search placeholder="今天想食点撒子喔～" v-model="foodName" @search="handlerSearch"></u-search>
		<div class=" list">
			<u-list @scrolltolower="scrolltolower">
				<u-list-item v-for="(item, index) in searchList" :key="index">
					<u-cell :title="item.title"></u-cell>
				</u-list-item>
			</u-list>
		</div>
	</view>
</template>

<script>
	import {
		searchFoods
	} from "@/request/api.js"
	export default {
		data() {
			return {
				foodName: "",
				searchList: []
			}
		},
		onLoad() {},
		methods: {
			scrolltolower() {
				console.log(1)
			},
			async handlerSearch(value) {
				const res = await searchFoods({
					foodName: value || ""
				})
				this.searchList = res.result.length === 0 ? this.searchList : res.result
			}
		}
	}
</script>

<style>

</style>
<template>
	<u-skeleton rows="10" :rowsHeight="20" :loading="result === null" :title="false">
		<div class="card">
			<!-- 品牌名称点击进入search -->
			<h2 class="food_name">{{result.food_name}}</h2>
			<u--text  type="primary"  :text="result.brand" @click="searchBrand"></u--text>
		</div>
		<div class="card">
			<h4>营养元素</h4>
			<u--text type="info" :text="result.server_size"></u--text>
			<div v-for="(value,key) in result.nutrient_details" style="display: flex;margin: 16px 0;" :key="value">
				<u--text :text="key"></u--text>
				<u--text type="info" :text="value"></u--text>
			</div>
		</div>
		<div class="card" v-for="(item,idx) in result?result.other_key_slice:[]" :key="idx">
			<h4>{{item}}</h4>
			<div  style="margin: 16px 0;"  v-for="(valueItem,valueIdx) in formatKeyValueSlice(result.other_value_slice[idx])" :key="valueIdx">
				<u--text type="primary"  :text="valueItem.label"
					@click="toggleCurrentDetails(valueItem)"></u--text>
			</div>
		</div>
		</view>
	</u-skeleton>
</template>

<script>
	// import mpvueEcharts from "@/components/echarts/echarts.vue"
	// import * as echarts from "@/components/echarts/echarts.js"
	import {
		linkDetails
	} from "@/request/api.js"

	export default {
		components: {
			// mpvueEcharts
		},
		data() {
			return {
				option: {
					backgroundColor: "#ffffff",
					series: [{
						label: {
							normal: {
								fontSize: 14
							}
						},
						type: 'pie',
						center: ['50%', '50%'],
						radius: ['0%', '40%'],
						data: [{
							value: 55,
							name: '北京'
						}, {
							value: 20,
							name: '武汉'
						}, {
							value: 10,
							name: '杭州'
						}, {
							value: 20,
							name: '广州'
						}, {
							value: 38,
							name: '上海'
						}]
					}]
				},
				result: null
			}
		},
		methods: {
			searchBrand(){
				uni.reLaunch({
					url:`/pages/index/index?foodName=${this.result.brand}`
				})
			},
			// 当前页面切换details
			toggleCurrentDetails(valueItem) {
				const {value} = valueItem
				uni.navigateTo({
					url: `/pages/details/details?link=${value}`,
				})
			},
			formatKeyValueSlice(arr) {
				const retArr = []
				for (let i = 0; i < arr.length; i += 2) {
					retArr.push({
						label: arr[i],
						value: arr[i + 1]
					})
				}
				return retArr
			},
			initChart(canvas, width, height) {
				chart = echarts.init(canvas, null, {
					width: width,
					height: height
				});
				canvas.setChart(chart);
				return chart;
			},
			async handleGetDetails({
				link
			}) {
				const res = await linkDetails({
					link
				})
				// console.log(link)
				// console.log(res.result)
				this.result = res.result
			}
		},
		onLoad(params) {
			this.handleGetDetails(params)
		}
	}
</script>

<style>
/* 	/deep/ .u-text{
		margin:18px 6px !important;
	} */
	.card {
		width: 90vw;
		margin: 10px auto;
		background-color: #fff;
		padding: 20px;
		box-sizing: border-box;
		border-radius: 12px;
		box-shadow: rgba(6, 24, 44, 0.4) 0px 0px 0px 2px, rgba(6, 24, 44, 0.65) 0px 4px 6px -1px, rgba(255, 255, 255, 0.08) 0px 1px 0px inset;
	}

	.food_name {
		font-size: 22px;
	}

	.ec-canvas {
		width: 100rpx;
		height: 100rpx;
	}
</style>
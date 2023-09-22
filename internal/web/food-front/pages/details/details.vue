<template>
	<u-skeleton rows="10" :rowsHeight="20" :loading="result === null" :title="false">
		<div class="card">
			<h2 class="food_name">{{result.food_name}}</h2>
			<u--text mode="link" :text="result.brand" href="https://www.uviewui.com"></u--text>
		</div>
		<div class="card">
			<h4>营养元素</h4>
			<u--text type="info" :text="result.server_size"></u--text>
			<div v-for="(value,key) in result.nutrient_details" style="display: flex;margin: 16px 0;" :key="value">
				<u--text :text="key"></u--text>
				<u--text type="info" :text="value"></u--text>
			</div>
		</div>
		<div class="card" v-for="(item,idx) in result.other_key_slice" :key="idx">
			<h4>{{item}}</h4>
			<div v-for="(valueItem,valueIdx) in formatKeyValueSlice(result.other_value_slice[idx])" :key="valueIdx">
				<u--text type="primary" :text="valueItem.label"></u--text>
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
	import f2 from "@antv/wx-f2"

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
	.card {
		width: 90vw;
		margin: 10px auto;
		background-color: #fff;
		padding: 20px;
		box-sizing: border-box;
		border-radius: 12px;
		box-shadow: rgba(0, 0, 0, .4) 0 2px 4px, rgba(0, 0, 0, .3) 0 7px 13px -3px, rgba(0, 0, 0, .2) 0 -3px 0 inset;
	}

	.food_name {
		font-size: 22px;
	}

	.ec-canvas {
		width: 100rpx;
		height: 100rpx;
	}
</style>
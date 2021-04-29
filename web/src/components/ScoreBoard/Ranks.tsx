import { ResponsiveBar } from '@nivo/bar'

import React, {useEffect} from "react";
import {SimpleReport} from "../../types/types";

const darkTheme = {
    axis: {
        fontSize: "14px",
        tickColor: "#eee",
        ticks: {
            line: {
                stroke: "#555555"
            },
            text: {
                fill: "#ffffff"
            }
        },
        legend: {
            text: {
                fill: "#aaaaaa"
            }
        }
    },
    grid: {
        line: {
            stroke: "#555555"
        }
    },
    tooltip: {
        container: {
            background: '#333',
        },
    },
};

type RanksProps = {
    report: SimpleReport
    isDarkTheme: boolean
}


export default function Ranks(props: RanksProps) {
    const report = props.report
    const data: Record<string, number | string> [] = []
    const dataKeys = new Set<string>();
    useEffect(() => {
        document.title = "Ranks"
    }, []);
    if ("Teams" in report){
        for (const team in report.Teams) {
            if (report.Teams.hasOwnProperty(team)) {
                const serviceAggregator: Record<string, number> = {}
                for (const host in report.Teams[team].Hosts){
                    if (report.Teams[team].Hosts.hasOwnProperty(host)) {
                        if (Object.keys(report.Teams[team].Hosts[host].Services).length !== 0){
                            for (const service in report.Teams[team].Hosts[host].Services) {
                                if (report.Teams[team].Hosts[host].Services.hasOwnProperty(service)) {
                                    const hst = report.Teams[team].Hosts[host]
                                    const sr = report.Teams[team].Hosts[host].Services[service]
                                    let keyName = ""
                                    if (sr.DisplayName){
                                        keyName = sr.DisplayName
                                    } else {
                                        if (hst.HostGroup !== undefined){
                                            keyName = hst.HostGroup.Name + "-" + sr.Name
                                        } else{
                                            keyName = sr.Name
                                        }
                                    }
                                    dataKeys.add(keyName)
                                    if (keyName in serviceAggregator){
                                        serviceAggregator[keyName] += sr.Points + sr.PointsBoost
                                    } else {
                                        serviceAggregator[keyName] = sr.Points + sr.PointsBoost
                                    }
                                }
                            }

                        }
                    }
                }

                if (Object.keys(serviceAggregator).length !== 0){
                    data.push({
                        ...serviceAggregator,
                        teamName: report.Teams[team].Name,
                    })
                }
            }
        }
    }

    const serviceSum = (teamObj: Record<string, string | number>) => {
        let sum: number = 0
        Object.keys(teamObj).forEach(function(key) {
            if (key !== "teamName"){
                sum += teamObj[key] as number;
            }
        });
        return sum
    }


    data.sort((a, b) => (serviceSum(a) === serviceSum(b)) ? (a.teamName > b.teamName ? 1 : -1) : (serviceSum(a) > serviceSum(b) ? 1 : -1))
    const theme = {fontSize: "0.875rem"}
    if (props.isDarkTheme){
        Object.assign(theme, darkTheme);
    }

    // @ts-ignore
    const TotalLabels = ({ bars, yScale }) => {
        // space between top of stacked bars and total label
        const labelMargin = 20;

        // @ts-ignore
        return bars.map(({ data: { data, indexValue }, x, width }, i) => {
            // sum of all the bar values in a stacked bar
            const total = Object.keys(data)
                // f ilter out whatever your indexBy value is
                .filter(key => key !== "teamName")
                .reduce((a, key) => a + data[key], 0);

            return (
                <g
                    transform={`translate(${x}, ${yScale(total) - labelMargin})`}
                    key={`${indexValue}-${i}`}
                >
                    <text
                        // add any class to the label here
                        className="bar-total-label"
                        x={width / 2}
                        y={labelMargin / 2}
                        textAnchor="middle"
                        alignmentBaseline="central"
                        // add any style to the label here
                        style={{
                            fill: props.isDarkTheme ? "rgb(222,255,255)" : "rgb(51,51,51)"
                        }}
                    >
                        {total}
                    </text>
                </g>
            );
        });
    };

    return (
        // @ts-ignore
        <ResponsiveBar
            theme={theme}
            data={data}
            keys={Array.from(dataKeys)}
            indexBy="teamName"
            margin={{ top: 50, right: 60, bottom: 50, left: 60 }}
            padding={0.3}
            colors={{ scheme: props.isDarkTheme ? 'nivo' : 'dark2' }}
            borderColor={{ from: 'color', modifiers: [ [ 'darker', '0' ] ] }}
            axisTop={null}
            axisRight={null}
            layers={["grid", "axes", "bars", TotalLabels, "markers", "legends"]}
            axisLeft={{
                tickSize: 5,
                tickPadding: 5,
                tickRotation: 0,
                legend: 'points',
                legendPosition: 'middle',
                legendOffset: -40
            }}
            labelSkipWidth={8}
            labelSkipHeight={12}
            labelTextColor={{ from: 'color', modifiers: [ [ 'darker', 2 ] ] }}
            legends={[]}
            animate={true}
            motionStiffness={70}
            motionDamping={15}
        />
    );
}
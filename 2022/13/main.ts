import { readFileSync } from "fs";

const getData = () => {
  let lines: string[];
  try {
    const data = readFileSync("input.txt", "utf8");
    lines = data.split("\r\n");
    return lines;
  } catch (error) {
    console.log("error", error);
  }
};

const compare = (packet1: number | any[], packet2: number | any[]): number => {
  if (typeof packet1 === "number" && typeof packet2 === "number") {
    return packet2 - packet1;
  } else if (Array.isArray(packet1) && typeof packet2 === "number") {
    return compare(packet1, [packet2]);
  } else if (typeof packet1 === "number" && Array.isArray(packet2)) {
    return compare([packet1], packet2);
  } else if (Array.isArray(packet1) && Array.isArray(packet2)) {
    for (let i = 0; i < packet1.length && i < packet2.length; i++) {
      const res = compare(packet1[i], packet2[i]);
      if (res === 0) {
        continue;
      }
      return res;
    }
    return packet2.length - packet1.length;
  }
  return 0;
};

const lines = getData() as string[];

const indices: number[] = [];
let index = 1;
for (let i = 0; i + 1 < lines.length; i += 3) {
  if (compare(JSON.parse(lines[i]), JSON.parse(lines[i + 1])) > 0) {
    indices.push(index);
  }
  index++;
}
// part 1
console.log(indices.reduce((partial, a) => partial + a, 0));

// part 2
const sortedLines = lines
  .filter((line) => line !== "")
  .sort((a, b) => compare(JSON.parse(b), JSON.parse(a)));

const d1Index = sortedLines.indexOf("[[2]]") + 1;
const d2Index = sortedLines.indexOf("[[6]]") + 1;
console.log(d1Index * d2Index);

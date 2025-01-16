import axios from "axios";
import React, { useState } from "react";
import "./Grid.css";

const Grid = () => {
     const [grid] = useState(Array(20).fill(Array(20).fill(null)));
     const [start, setStart] = useState(null);
     const [end, setEnd] = useState([]);
     const [path, setPath] = useState([]);

     const handleCellClick = (x, y) => {
          if (!start) {
               setStart({ x, y });
          } else if (!end) {
               setEnd({ x, y });
          }
     };

     const calculatePath = async () => {
          console.log(start, "start", end, "end");
          //   if (start && end) {
          const response = await axios.post("https://localhost:8080/find-path", {
               start,
               end,
          });
          setPath(response.data.path);
          //   }
     };

     const resetGrid = () => {
          setStart(null);
          setEnd(null);
          setPath([]);
     };

     return (
          <div>
               <div className="grid">
                    {grid.map((row, rowIndex) => {
                         row.map((_, colIndex) => {
                              const isStart = start?.x === rowIndex && start?.y === colIndex;
                              const isEnd = end?.x === rowIndex && end?.y === colIndex;
                              const isPath = path.some((point) => point.x === rowIndex && point.y === colIndex);

                              return (
                                   <div
                                        key={`${rowIndex}-${colIndex}`}
                                        className={`cell ${isStart ? "start" : ""}
                                                ${isEnd ? "end" : ""} ${isPath ? "path" : ""}
                                        `}
                                        onClick={() => handleCellClick(rowIndex, colIndex)}
                                   ></div>
                              );
                         });
                    })}
               </div>
               <button onClick={calculatePath}>Calculate Path</button>
               <button onClick={resetGrid}>Reset Grid</button>
          </div>
     );
};

export default Grid;

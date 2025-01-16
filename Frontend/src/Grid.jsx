import axios from "axios";
import React, { useState } from "react";
import "./Grid.css";

const Grid = () => {
     const [grid] = useState(
          Array(20)
               .fill()
               .map(() => Array(20).fill(null))
     );
     const [start, setStart] = useState(null);
     const [end, setEnd] = useState(null);
     const [path, setPath] = useState([]);

     const handleCellClick = (x, y) => {
          if (!start) {
               setStart({ x, y });
          } else if (!end) {
               setEnd({ x, y });
          }
     };

     const calculatePath = async () => {
          try {
               if (!start || !end) {
                    alert("Please select both start and end points");
                    return;
               }

               const response = await axios.post("http://localhost:8080/find-path", {
                    start,
                    end,
               });

               if (response.data.error) {
                    alert(response.data.error);
                    return;
               }

               setPath(response.data.path || []);
          } catch (error) {
               console.error("Error calculating path:", error);
               alert("Error calculating path. Please try again.");
          }
     };

     const resetGrid = () => {
          setStart(null);
          setEnd(null);
          setPath([]);
     };

     return (
          <div className="container">
               <h1>Path Finder</h1>
               <div className="grid">
                    {grid.map((row, rowIndex) => (
                         <div key={rowIndex} className="row">
                              {row.map((_, colIndex) => {
                                   const isStart = start?.x === rowIndex && start?.y === colIndex;
                                   const isEnd = end?.x === rowIndex && end?.y === colIndex;
                                   const isPath = path.some((point) => point.x === rowIndex && point.y === colIndex);

                                   return (
                                        <div
                                             key={`${rowIndex}-${colIndex}`}
                                             className={`cell ${isStart ? "start" : ""} 
                                        ${isEnd ? "end" : ""} ${isPath ? "path" : ""}`}
                                             onClick={() => handleCellClick(rowIndex, colIndex)}
                                        />
                                   );
                              })}
                         </div>
                    ))}
               </div>
               <div className="buttons">
                    <button onClick={calculatePath}>Calculate Path</button>
                    <button onClick={resetGrid}>Reset Grid</button>
               </div>
          </div>
     );
};

export default Grid;

import { useState } from "react";
const colors = ["white", "blue", "black", "red", "green"];
const landTypes = [
  "fetch",
  "tango",
  "shock",
  "triome",
  "surveil",
  "verge",
  "bond",
  "pain",
  "horizon",
  "check",
  "slow",
  "gates",
  "thriving",
  "rainbow",
];

function ManaForm() {
  const [selectedColors, setSelectedColors] = useState<string[]>(colors);
  const [selectedLandTypes, setSelectedLandTypes] =
    useState<string[]>(landTypes);

  const toggleAllItems = (
    items: string[],
    selected: string[],
    setSelected: any,
  ) => {
    if (selected.length === items.length) {
      setSelected([]);
    } else {
      setSelected(items);
    }
  };

  const toggleItem = (item: string, selected: string[], setSelected: any) => {
    if (selected.includes(item)) {
      setSelected(selected.filter((i) => i !== item));
    } else {
      setSelected([...selected, item]);
    }
  };

  const allChecked = (items: string[], selected: string[]) => {
    return items.length === selected.length;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const request = {
      colors: selectedColors,
      landTypes: selectedLandTypes,
    };
    try {
      const res = await fetch("http://localhost:8080/api/lands", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(request),
      });
      const data = await res.json();
      console.log(data);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <form className="flex flex-col gap-8 max-w-md">
      {/* Color Identity Section */}
      <div>
        <div className="flex gap-2 items-center mb-2">
          <h2 className="text-lg font-semibold">Color Identity</h2>
          <label className="flex items-center gap-1 text-sm">
            <input
              type="checkbox"
              checked={allChecked(colors, selectedColors)}
              onChange={() =>
                toggleAllItems(colors, selectedColors, setSelectedColors)
              }
            />
            Check All
          </label>
        </div>
        <div className="flex flex-wrap gap-2">
          {colors.map((color) => (
            <label
              key={color}
              className="flex items-center gap-2 px-3 py-1 border border-gray-300 rounded-md cursor-pointer hover:bg-gray-500"
            >
              <input
                type="checkbox"
                name="colors"
                value={color}
                className="cursor-pointer"
                checked={selectedColors.includes(color)}
                onChange={() =>
                  toggleItem(color, selectedColors, setSelectedColors)
                }
              />
              <span className="capitalize">{color}</span>
            </label>
          ))}
        </div>
      </div>

      {/* Land Types Section */}
      <div>
        <div className="flex gap-2 items-center mb-2">
          <h2 className="text-lg font-semibold mb-2 text-left">Land Types</h2>
          <label className="flex items-center gap-1 text-sm">
            <input
              type="checkbox"
              checked={allChecked(landTypes, selectedLandTypes)}
              onChange={() =>
                toggleAllItems(
                  landTypes,
                  selectedLandTypes,
                  setSelectedLandTypes,
                )
              }
            />
            Check All
          </label>
        </div>
        <div className="flex flex-wrap gap-2">
          {landTypes.map((type) => (
            <label
              key={type}
              className="flex items-center gap-2 px-3 py-1 border border-gray-300 rounded-md cursor-pointer hover:bg-gray-500"
            >
              <input
                type="checkbox"
                name="landTypes"
                value={type}
                className=" cursor-pointer"
                checked={selectedLandTypes.includes(type)}
                onChange={() =>
                  toggleItem(type, selectedLandTypes, setSelectedLandTypes)
                }
              />
              <span className="capitalize">{type}</span>
            </label>
          ))}
        </div>
      </div>
      <button onClick={handleSubmit}>Generate Lands</button>
    </form>
  );
}

export default ManaForm;

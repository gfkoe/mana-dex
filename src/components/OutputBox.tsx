function OutputBox({ text }: { text: string }) {
  return (
    <textarea
      className="w-full h-full flex flex-col gap-8 max-w-md bg-white text-black"
      value={text}
      readOnly
    />
  );
}
export default OutputBox;

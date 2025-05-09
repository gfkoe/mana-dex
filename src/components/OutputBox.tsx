function OutputBox({ text }: { text: string }) {
  const copyToClipboard = () => {
    navigator.clipboard.writeText(text);
  };
  return (
    <>
      <textarea
        className="w-full h-full flex flex-col gap-8 max-w-md bg-white text-black"
        value={text}
        readOnly
      />
      <button onClick={copyToClipboard}>Copy to Clipboard</button>
    </>
  );
}
export default OutputBox;

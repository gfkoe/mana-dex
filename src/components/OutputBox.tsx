function OutputBox({ text }: { text: string }) {
  return <textarea className="bg-white text-black" value={text} readOnly />;
}
export default OutputBox;

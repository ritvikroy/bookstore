import { tokensAtom } from "../utils/atoms";
import { useAtom } from "jotai";
const useAtomsApi = () => {
  const [tokensData, setTokensData] = useAtom(tokensAtom);

  return {
    tokensData,
    setTokensData,
  };
};

export default useAtomsApi;

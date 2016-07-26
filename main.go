package main

//import "github.com/davecgh/go-spew/spew"
import (
	"./alignment"
	a "./types/amino"
	n "./types/nucleic"
	"fmt"
	"github.com/davecheney/profile"
	"time"
)

func test() *alignment.Alignment {
	cmp := alignment.NewAlignment(
		//[]n.NucleicAcid{n.G, n.C, n.G, n.A, n.T, n.G, n.A, n.G, n.C, n.T, n.T, n.G},
		//[]a.AminoAcid{a.M, a.R, a.A, a.C},
		//[]n.NucleicAcid{n.C, n.C, n.T, n.C, n.A, n.A, n.A, n.T, n.C, n.T, n.A, n.G, n.T, n.C, n.T, n.T, n.T, n.G, n.G, n.C},
		//[]a.AminoAcid{a.P, a.Q, a.I, a.T, a.L, a.W},
		n.ReadString("CCTCAAATCACTCTTTGGCAACGACCCATCGTCACAATAAAGATAGGGGGGCAGCTAARGGAAGCTCTATTAGATACAGGAGCAGATGATACAGTATTAGAAGATATAAATTTGCCAGGAAGATGGACACCAAAAATKATAGTGGGAATTGGAGGTTTTACCAAAGTAAGACAGTATGATCAGATACCTGTAGAAATTTGTGGACATAAAGCTATAGGTACAGTRTTAGTAGGACCTACACCTGCCAACATAATTGGAAGAAATCTGTTGACYCAGATTGGTTGCACTTTAAATTTTCCCATTAGTCCTATTGACACTGTACCAGTAAAATTAAAGCCAGGAATGGATGGCCCAAAAGTTAAACAATGGCCATTGACAGAAGAAAAAATAAAAGCATTAGTAGAAATTTGTGCAGAATTGGAASAGGACGGGAAAATTTCAAAAATTGGGCCTGAAAATCCATACAATACTCCAGTATTTGCCATAAAGAAAAAGAACAGYGATAAATGGAGAAAATTAGTAGATTTCAGAGAACTTAATAAGAGAACTCAAGACTTCTGGGAAGTTCAATTAGGAATACCACATCCCGGAGGGTTAAAAAAGAACAAATCAGTAACAGTACTGGATGTGGGTGATGCATATTTTTCARTTCCCTTAGATGAAGACTTCAGGAAGTATACTGCATTTACCATACCTAGTATAAACAATGAGACACCAGGGACTAGATATCAGTACAATGTGCTTCCACAGGGATGGAAAGGATCACCAGCAATATTCCAAAGTAGCATGACAAGAATCTTAGAACCTTTTAGAAAACAGAATCCAGACATAGTTATCTGTCAATAYGTGGATGATTTGTATGTAGGATCTGACTTAGAAATAGAGMAGCATAGAACAAAAGTAGAGGAACTGAGACAACATTTGTGGAAGTGGGGNTTTTACACACCAGACAAMAAACATCAGAAAGAACCTCCATTCCTTTGGATGGGTTATGAACTCCATCCTGATAAATGGACA"),
		//a.ReadString("PQITLWQRPLVTIKIGGQLKEALLDTGADDTVLEEMNLPGRWKPKMIGGIGGFIKVRQYDQILIEICGHKAIGTVLVGPTPVNIIGRNLLTQIGCTLNF"),
		a.ReadString("PISPIETVPVKLKPGMDGPKVKQWPLTEEKIKALVEICTEMEKEGKISKIGPENPYNTPVFAIKKKDSTKWRKLVDFRELNKRTQDFWEVQLGIPHPAGLKKKKSVTVLDVGDAYFSVPLDKDFRKYTAFTIPSINNETPGIRYQYNVLPQGWKGSPAIFQSSMTKILEPFRKQNPDIVIYQYMDDLYVGSDLEIGQHRTKIEELRQHLLRWGFTTPDKKHQKEPPFLWMGYELHPDKWTVQPIVLPEKDSWTVNDIQKLVGKLNWASQIYAGIKVKQLCKLLRGTKALTEVIPLTEEAELELAENREILKEPVHGVYYDPSKDLIAEIQKQGQGQWTYQIYQEPFKNLKTGKYARMRGAHTNDVKQLTEAVQKIATESIVIWGKTPKFKLPIQKETWEAWWTEYWQATWIPEWEFVNTPPLVKLWYQLEKEPIVGAETFYVDGAANRETKLGKAGYVTDRGRQKVVSLTDTTNQKTELQAIHLALQDSGLEVNIVTDSQYALGIIQAQPDKSESELVSQIIEQLIKKEKVYLAWVPAHKGIGGNEQVDKLVSAGIRKVL"),
		10,
		2,
	)
	cmp.CalcScore()
	return cmp
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	var cmp *alignment.Alignment
	start := time.Now()
	for i := 0; i < 10; i++ {
		cmp = test()
	}
	elapsed := time.Since(start)
	fmt.Printf("Binomial took %s\n", elapsed)
	fmt.Printf("Calc times %d\n", cmp.GetCalcTimes())
	//fmt.Printf("%s", cmp.GetScorePath())
	//table := alignment.GetMutationScoreTable()
	//spew.Print(table)
}

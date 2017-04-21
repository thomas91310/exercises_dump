package main

import (
	"fmt"
	"strconv"
)

const (
	hugeNumber = "37107287533902102798797998220837590246510135740250463769376774900097126481248969700780504170182605387432498619952474105947423330951305812372661730962991942213363574161572522430563301811072406154908250230675882075393461711719803104210475137780632466768926167069662363382013637841838368417873436172675728112879812849979408065481931592621691275889832738442742289174325203219235894228767964876702721893184745144573600130643909116721685684458871160315327670386486105843025439939619828917593665686757934951621764571418565606295021572231965867550793241933316490635246274190492910143244581382266334794475817892575867718337217661963751590579239728245598838407582035653253593990084026335689488301894586282278288018119938482628201427819413994056758715117009439035398664372827112653829987240784473053190104293586865155060062958648615320752733719591914205172558297169388870771546649911559348760353292171497005693854370070576826684624621495650076471787294438377604532826541087568284431911906346940378552177792951453612327252500029607107508256381565671088525835072145876576172410976447339110607218265236877223636045174237069058518606604482076212098132878607339694128114266041808683061932846081119106155694051268969251934325451728388641918047049293215058642563049483624672216484350762017279180399446930047329563406911573244438690812579451408905770622942919710792820955037687525678773091862540744969844508330393682126183363848253301546861961243487676812975343759465158038628759287849020152168555482871720121925776695478182833757993103614740356856449095527097864797581167263201004368978425535399209318374414978068609844840309812907779179908821879532736447567559084803087086987551392711854517078544161852424320693150332599594068957565367821070749269665376763262354472106979395067965269474259770973916669376304263398708541052684708299085211399427365734116182760315001271653786073615010808570091499395125570281987460043753582903531743471732693212357815498262974255273730794953759765105305946966067683156574377167401875275889028025717332296191766687138199318110487701902712526768027607800301367868099252546340106163286652636270218540497705585629946580636237993140746255962240744869082311749777923654662572469233228109171419143028819710328859780666976089293863828502533340334413065578016127815921815005561868836468420090470230530811728164304876237919698424872550366387845831148769693215490281042402013833512446218144177347063783299490636259666498587618221225225512486764533677201869716985443124195724099139590089523100588229554825530026352078153229679624948164195386821877476085327132285723110424803456124867697064507995236377742425354112916842768655389262050249103265729672370191327572567528565324825826546309220705859652229798860272258331913126375147341994889534765745501184957014548792889848568277260777137214037988797153829820378303147352772158034814451349137322665138134829543829199918180278916522431027392251122869539409579530664052326325380441000596549391598795936352974615218550237130764225512118369380358038858490341698116222072977186158236678424689157993532961922624679571944012690438771072750481023908955235974572318970677254791506150550495392297953090112996751986188088225875314529584099251203829009407770775672113067397083047244838165338735023408456470580773088295174767140363198008187129011875491310547126581976233310448183862695154563349263665728975634005004284628018351707052783183942588214552122725125032755121603546981200581762165212827652751691296897789322381957343293399464375019078369457658833523998867550616496518477518073816883786109152735792970133761680449839931733127313292418570714734956691667468763466091503591467750499518671430235219628894890102423325116913619626622732674608005915474718307983928685352069469445407247684182252467441716151403642798227334805555621481897142617910342598647204516893989422179826088076852877836461827993463137677543078093633330189826420901084880252167467088321512018588354322381287695278671329612474782464538636993009049310363619763878039621840735723997942234062353938083396513274080111166662789198148808779794187687614423003098449085141160661826293682836764744779239180335110989069790714857869440895529906536404474255760836599766457950966602439640990538960712019821997604759949019723029764913982680032973156037120041377903785566085089252167309393198727502754689069037075394130426523150119480937724504879515095410092164586375471059843679178639167021187492431995700641917969777599028300699153687137119366149528113058763802784107544497330784078992311553556256114232242325503368544248891735344889911501440648020369068063960672322193204149535415031288803395360532993403680069777106505666319548123488067321014673905856855793458140362782270328082616570773948327592232845941706525094512325230608229188020587731971983945018088807242966198081119777158542502016545090413245809786882778948721859617721078384350691861554356628840622574736922845095162084960398013400172393067166682355524525280460972253503534226472524250874054075591789781264330331690"
)

func sumFirst10Digits(number string) (int, error) {
	sum := 0
	for _, c := range number[:10] {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			return 0, fmt.Errorf("the number passed as the first number doesn't contain just numbers : %v", err)
		}
		sum += n
	}
	return sum, nil
}

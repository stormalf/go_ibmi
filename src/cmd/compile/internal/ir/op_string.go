// Code generated by "stringer -type=Op -trimprefix=O node.go"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OLITERAL-4]
	_ = x[ONIL-5]
	_ = x[OADD-6]
	_ = x[OSUB-7]
	_ = x[OOR-8]
	_ = x[OXOR-9]
	_ = x[OADDSTR-10]
	_ = x[OADDR-11]
	_ = x[OANDAND-12]
	_ = x[OAPPEND-13]
	_ = x[OBYTES2STR-14]
	_ = x[OBYTES2STRTMP-15]
	_ = x[ORUNES2STR-16]
	_ = x[OSTR2BYTES-17]
	_ = x[OSTR2BYTESTMP-18]
	_ = x[OSTR2RUNES-19]
	_ = x[OSLICE2ARR-20]
	_ = x[OSLICE2ARRPTR-21]
	_ = x[OAS-22]
	_ = x[OAS2-23]
	_ = x[OAS2DOTTYPE-24]
	_ = x[OAS2FUNC-25]
	_ = x[OAS2MAPR-26]
	_ = x[OAS2RECV-27]
	_ = x[OASOP-28]
	_ = x[OCALL-29]
	_ = x[OCALLFUNC-30]
	_ = x[OCALLMETH-31]
	_ = x[OCALLINTER-32]
	_ = x[OCAP-33]
	_ = x[OCLEAR-34]
	_ = x[OCLOSE-35]
	_ = x[OCLOSURE-36]
	_ = x[OCOMPLIT-37]
	_ = x[OMAPLIT-38]
	_ = x[OSTRUCTLIT-39]
	_ = x[OARRAYLIT-40]
	_ = x[OSLICELIT-41]
	_ = x[OPTRLIT-42]
	_ = x[OCONV-43]
	_ = x[OCONVIFACE-44]
	_ = x[OCONVNOP-45]
	_ = x[OCOPY-46]
	_ = x[ODCL-47]
	_ = x[ODCLFUNC-48]
	_ = x[ODELETE-49]
	_ = x[ODOT-50]
	_ = x[ODOTPTR-51]
	_ = x[ODOTMETH-52]
	_ = x[ODOTINTER-53]
	_ = x[OXDOT-54]
	_ = x[ODOTTYPE-55]
	_ = x[ODOTTYPE2-56]
	_ = x[OEQ-57]
	_ = x[ONE-58]
	_ = x[OLT-59]
	_ = x[OLE-60]
	_ = x[OGE-61]
	_ = x[OGT-62]
	_ = x[ODEREF-63]
	_ = x[OINDEX-64]
	_ = x[OINDEXMAP-65]
	_ = x[OKEY-66]
	_ = x[OSTRUCTKEY-67]
	_ = x[OLEN-68]
	_ = x[OMAKE-69]
	_ = x[OMAKECHAN-70]
	_ = x[OMAKEMAP-71]
	_ = x[OMAKESLICE-72]
	_ = x[OMAKESLICECOPY-73]
	_ = x[OMUL-74]
	_ = x[ODIV-75]
	_ = x[OMOD-76]
	_ = x[OLSH-77]
	_ = x[ORSH-78]
	_ = x[OAND-79]
	_ = x[OANDNOT-80]
	_ = x[ONEW-81]
	_ = x[ONOT-82]
	_ = x[OBITNOT-83]
	_ = x[OPLUS-84]
	_ = x[ONEG-85]
	_ = x[OOROR-86]
	_ = x[OPANIC-87]
	_ = x[OPRINT-88]
	_ = x[OPRINTLN-89]
	_ = x[OPAREN-90]
	_ = x[OSEND-91]
	_ = x[OSLICE-92]
	_ = x[OSLICEARR-93]
	_ = x[OSLICESTR-94]
	_ = x[OSLICE3-95]
	_ = x[OSLICE3ARR-96]
	_ = x[OSLICEHEADER-97]
	_ = x[OSTRINGHEADER-98]
	_ = x[ORECOVER-99]
	_ = x[ORECOVERFP-100]
	_ = x[ORECV-101]
	_ = x[ORUNESTR-102]
	_ = x[OSELRECV2-103]
	_ = x[OMIN-104]
	_ = x[OMAX-105]
	_ = x[OREAL-106]
	_ = x[OIMAG-107]
	_ = x[OCOMPLEX-108]
	_ = x[OUNSAFEADD-109]
	_ = x[OUNSAFESLICE-110]
	_ = x[OUNSAFESLICEDATA-111]
	_ = x[OUNSAFESTRING-112]
	_ = x[OUNSAFESTRINGDATA-113]
	_ = x[OMETHEXPR-114]
	_ = x[OMETHVALUE-115]
	_ = x[OBLOCK-116]
	_ = x[OBREAK-117]
	_ = x[OCASE-118]
	_ = x[OCONTINUE-119]
	_ = x[ODEFER-120]
	_ = x[OFALL-121]
	_ = x[OFOR-122]
	_ = x[OGOTO-123]
	_ = x[OIF-124]
	_ = x[OLABEL-125]
	_ = x[OGO-126]
	_ = x[ORANGE-127]
	_ = x[ORETURN-128]
	_ = x[OSELECT-129]
	_ = x[OSWITCH-130]
	_ = x[OTYPESW-131]
	_ = x[OINLCALL-132]
	_ = x[OMAKEFACE-133]
	_ = x[OITAB-134]
	_ = x[OIDATA-135]
	_ = x[OSPTR-136]
	_ = x[OCFUNC-137]
	_ = x[OCHECKNIL-138]
	_ = x[ORESULT-139]
	_ = x[OINLMARK-140]
	_ = x[OLINKSYMOFFSET-141]
	_ = x[OJUMPTABLE-142]
	_ = x[OINTERFACESWITCH-143]
	_ = x[ODYNAMICDOTTYPE-144]
	_ = x[ODYNAMICDOTTYPE2-145]
	_ = x[ODYNAMICTYPE-146]
	_ = x[OTAILCALL-147]
	_ = x[OGETG-148]
	_ = x[OGETCALLERSP-149]
	_ = x[OEND-150]
}

const _Op_name = "XXXNAMENONAMETYPELITERALNILADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESSLICE2ARRSLICE2ARRPTRASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCAPCLEARCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVNOPCOPYDCLDCLFUNCDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNOTBITNOTPLUSNEGORORPANICPRINTPRINTLNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERSTRINGHEADERRECOVERRECOVERFPRECVRUNESTRSELRECV2MINMAXREALIMAGCOMPLEXUNSAFEADDUNSAFESLICEUNSAFESLICEDATAUNSAFESTRINGUNSAFESTRINGDATAMETHEXPRMETHVALUEBLOCKBREAKCASECONTINUEDEFERFALLFORGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWINLCALLMAKEFACEITABIDATASPTRCFUNCCHECKNILRESULTINLMARKLINKSYMOFFSETJUMPTABLEINTERFACESWITCHDYNAMICDOTTYPEDYNAMICDOTTYPE2DYNAMICTYPETAILCALLGETGGETCALLERSPEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 24, 27, 30, 33, 35, 38, 44, 48, 54, 60, 69, 81, 90, 99, 111, 120, 129, 141, 143, 146, 156, 163, 170, 177, 181, 185, 193, 201, 210, 213, 218, 223, 230, 237, 243, 252, 260, 268, 274, 278, 287, 294, 298, 301, 308, 314, 317, 323, 330, 338, 342, 349, 357, 359, 361, 363, 365, 367, 369, 374, 379, 387, 390, 399, 402, 406, 414, 421, 430, 443, 446, 449, 452, 455, 458, 461, 467, 470, 473, 479, 483, 486, 490, 495, 500, 507, 512, 516, 521, 529, 537, 543, 552, 563, 575, 582, 591, 595, 602, 610, 613, 616, 620, 624, 631, 640, 651, 666, 678, 694, 702, 711, 716, 721, 725, 733, 738, 742, 745, 749, 751, 756, 758, 763, 769, 775, 781, 787, 794, 802, 806, 811, 815, 820, 828, 834, 841, 854, 863, 878, 892, 907, 918, 926, 930, 941, 944}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
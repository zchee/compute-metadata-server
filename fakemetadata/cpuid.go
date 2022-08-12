// Copyright 2022 The compute-metadata-server Authors
// SPDX-License-Identifier: BSD-3-Clause

package fakemetadata

import (
	cpuid "github.com/klauspost/cpuid/v2"
)

type X86Microarchitecture int

const (
	X86_UNKNOWN      X86Microarchitecture = iota
	INTEL_80486                           // https://en.wikichip.org/wiki/intel/microarchitectures/80486
	INTEL_P5                              // https://en.wikichip.org/wiki/intel/microarchitectures/p5
	INTEL_LAKEMONT                        // https://en.wikichip.org/wiki/intel/quark
	INTEL_CORE                            // https://en.wikipedia.org/wiki/Intel_Core_(microarchitecture)
	INTEL_PNR                             // https://en.wikipedia.org/wiki/Penryn_(microarchitecture)
	INTEL_NHM                             // https://en.wikipedia.org/wiki/Nehalem_(microarchitecture)
	INTEL_ATOM_BNL                        // https://en.wikipedia.org/wiki/Bonnell_(microarchitecture)
	INTEL_WSM                             // https://en.wikipedia.org/wiki/Westmere_(microarchitecture)
	INTEL_SNB                             // https://en.wikipedia.org/wiki/Sandy_Bridge#Models_and_steppings
	INTEL_IVB                             // https://en.wikipedia.org/wiki/Ivy_Bridge_(microarchitecture)#Models_and_steppings
	INTEL_ATOM_SMT                        // https://en.wikipedia.org/wiki/Silvermont
	INTEL_HSW                             // https://en.wikipedia.org/wiki/Haswell_(microarchitecture)
	INTEL_BDW                             // https://en.wikipedia.org/wiki/Broadwell_(microarchitecture)
	INTEL_SKL                             // https://en.wikipedia.org/wiki/Skylake_(microarchitecture)
	INTEL_ATOM_GMT                        // https://en.wikipedia.org/wiki/Goldmont
	INTEL_KBL                             // https://en.wikipedia.org/wiki/Kaby_Lake
	INTEL_CFL                             // https://en.wikipedia.org/wiki/Coffee_Lake
	INTEL_WHL                             // https://en.wikipedia.org/wiki/Whiskey_Lake_(microarchitecture)
	INTEL_CML                             // https://en.wikichip.org/wiki/intel/microarchitectures/comet_lake
	INTEL_CNL                             // https://en.wikipedia.org/wiki/Cannon_Lake_(microarchitecture)
	INTEL_ICL                             // https://en.wikipedia.org/wiki/Ice_Lake_(microprocessor)
	INTEL_TGL                             // https://en.wikipedia.org/wiki/Tiger_Lake_(microarchitecture)
	INTEL_SPR                             // https://en.wikipedia.org/wiki/Sapphire_Rapids
	INTEL_ADL                             // https://en.wikichip.org/wiki/intel/microarchitectures/alder_lake
	INTEL_RCL                             // https://en.wikichip.org/wiki/intel/microarchitectures/rocket_lake
	INTEL_KNIGHTS_M                       // https://en.wikichip.org/wiki/intel/microarchitectures/knights_mill
	INTEL_KNIGHTS_L                       // https://en.wikichip.org/wiki/intel/microarchitectures/knights_landing
	INTEL_KNIGHTS_F                       // https://en.wikichip.org/wiki/intel/microarchitectures/knights_ferry
	INTEL_KNIGHTS_C                       // https://en.wikichip.org/wiki/intel/microarchitectures/knights_corner
	INTEL_NETBURST                        // https://en.wikichip.org/wiki/intel/microarchitectures/netburst
	AMD_HAMMER                            // K8 HAMMER
	AMD_K10                               // K10
	AMD_K11                               // http://developer.amd.com/wordpress/media/2012/10/41788.pdf
	AMD_K12                               // https://www.amd.com/system/files/TechDocs/44739_12h_Rev_Gd.pdf
	AMD_BOBCAT                            // https://www.amd.com/system/files/TechDocs/47534_14h_Mod_00h-0Fh_Rev_Guide.pdf
	AMD_PILEDRIVER                        // https://en.wikichip.org/wiki/amd/microarchitectures/piledriver
	AMD_STREAMROLLER                      // https://en.wikichip.org/wiki/amd/microarchitectures/steamroller
	AMD_EXCAVATOR                         // https://en.wikichip.org/wiki/amd/microarchitectures/excavator
	AMD_BULLDOZER                         // https://en.wikichip.org/wiki/amd/microarchitectures/bulldozer
	AMD_JAGUAR                            // K16 JAGUAR
	AMD_PUMA                              // K16 PUMA
	AMD_ZEN                               // https://en.wikichip.org/wiki/amd/microarchitectures/zen
	AMD_ZEN_PLUS                          // https://en.wikichip.org/wiki/amd/microarchitectures/zen%2B
	AMD_ZEN2                              // https://en.wikichip.org/wiki/amd/microarchitectures/zen_2
	AMD_ZEN3                              // https://en.wikichip.org/wiki/amd/microarchitectures/zen_3
	AMD_ZEN4                              // https://en.wikichip.org/wiki/amd/microarchitectures/zen_4
)

func (x86 X86Microarchitecture) String() string {
	switch x86 {
	case X86_UNKNOWN:
		return "UNKNOWN"
	case INTEL_80486:
		return "Intel 80486"
	case INTEL_P5:
		return "Intel P5"
	case INTEL_LAKEMONT:
		return "Intel lakemont"
	case INTEL_CORE:
		return "Intel Core"
	case INTEL_PNR:
		return "Intel Penryn"
	case INTEL_NHM:
		return "Intel Nehalem"
	case INTEL_ATOM_BNL:
		return "Intel Bonnell"
	case INTEL_WSM:
		return "Intel Westmere"
	case INTEL_SNB:
		return "Intel Sandy Bridge"
	case INTEL_IVB:
		return "Intel Ivy Bridge"
	case INTEL_ATOM_SMT:
		return "Intel silvermont"
	case INTEL_HSW:
		return "Intel Haswell"
	case INTEL_BDW:
		return "Intel Broadwell"
	case INTEL_SKL:
		return "Intel Skylake"
	case INTEL_ATOM_GMT:
		return "Intel Goldmont"
	case INTEL_KBL:
		return "Intel Kaby Lake"
	case INTEL_CFL:
		return "Intel Coffee Lake"
	case INTEL_WHL:
		return "Intel Whiskey Lake"
	case INTEL_CML:
		return "Intel Comet Lake"
	case INTEL_CNL:
		return "Intel Cannon Lake"
	case INTEL_ICL:
		return "Intel Ice Lake"
	case INTEL_TGL:
		return "Intel Tiger Lake"
	case INTEL_SPR:
		return "Intel Sapphire Rapids"
	case INTEL_ADL:
		return "Intel Alder Lake"
	case INTEL_RCL:
		return "Intel Rocket Lake"
	case INTEL_KNIGHTS_M:
		return "Intel Knights Mill"
	case INTEL_KNIGHTS_L:
		return "Intel Knights Landing"
	case INTEL_KNIGHTS_F:
		return "Intel Knights Ferry"
	case INTEL_KNIGHTS_C:
		return "Intel Knights Corner"
	case INTEL_NETBURST:
		return "Intel Netburst"
	case AMD_HAMMER:
		return "AMD K8 Hammer"
	case AMD_K10:
		return "AMD K10"
	case AMD_K11:
		return "AMD K11"
	case AMD_K12:
		return "AMD K12"
	case AMD_BOBCAT:
		return "AMD K14 Bobcat"
	case AMD_PILEDRIVER:
		return "AMD K15 Piledriver"
	case AMD_STREAMROLLER:
		return "AMD K15 Streamroller"
	case AMD_EXCAVATOR:
		return "AMD K15 Excavator"
	case AMD_BULLDOZER:
		return "AMD K15 Bulldozer"
	case AMD_JAGUAR:
		return "AMD K16 Jaguar"
	case AMD_PUMA:
		return "AMD K16 Puma"
	case AMD_ZEN:
		return "AMD K17 Zen"
	case AMD_ZEN_PLUS:
		return "AMD K17 Zen+"
	case AMD_ZEN2:
		return "AMD K17 Zen 2"
	case AMD_ZEN3:
		return "AMD K19 Zen 3"
	case AMD_ZEN4:
		return "AMD K19 Zen 4"
	}

	return ""
}

func matchFamilyModel(info cpuid.CPUInfo, family, model int) bool {
	return info.Family == family && info.Model == model
}

func detectCPUMicroarchitecture(info cpuid.CPUInfo) X86Microarchitecture {
	switch info.VendorID {
	case cpuid.Intel:
		switch {
		case matchFamilyModel(info, 0x04, 0x01),
			matchFamilyModel(info, 0x04, 0x02),
			matchFamilyModel(info, 0x04, 0x03),
			matchFamilyModel(info, 0x04, 0x04),
			matchFamilyModel(info, 0x04, 0x05),
			matchFamilyModel(info, 0x04, 0x07),
			matchFamilyModel(info, 0x04, 0x08),
			matchFamilyModel(info, 0x04, 0x09):
			return INTEL_80486

		case matchFamilyModel(info, 0x05, 0x01),
			matchFamilyModel(info, 0x05, 0x02),
			matchFamilyModel(info, 0x05, 0x04),
			matchFamilyModel(info, 0x05, 0x07),
			matchFamilyModel(info, 0x05, 0x08):
			return INTEL_P5

		case matchFamilyModel(info, 0x05, 0x09),
			matchFamilyModel(info, 0x05, 0x0A):
			return INTEL_LAKEMONT

		case matchFamilyModel(info, 0x06, 0x1C), // Intel(R) Atom(TM) CPU 230 @ 1.60GHz
			matchFamilyModel(info, 0x06, 0x35),
			matchFamilyModel(info, 0x06, 0x36),
			matchFamilyModel(info, 0x06, 0x70): // https://en.wikichip.org/wiki/intel/atom/230
			return INTEL_ATOM_BNL

		case matchFamilyModel(info, 0x06, 0x37),
			matchFamilyModel(info, 0x06, 0x4C):
			return INTEL_ATOM_SMT

		case matchFamilyModel(info, 0x06, 0x5C):
			return INTEL_ATOM_GMT

		case matchFamilyModel(info, 0x06, 0x0F),
			matchFamilyModel(info, 0x06, 0x16):
			return INTEL_CORE

		case matchFamilyModel(info, 0x06, 0x17),
			matchFamilyModel(info, 0x06, 0x1D):
			return INTEL_PNR

		case matchFamilyModel(info, 0x06, 0x1A),
			matchFamilyModel(info, 0x06, 0x1E),
			matchFamilyModel(info, 0x06, 0x1F),
			matchFamilyModel(info, 0x06, 0x2E):
			return INTEL_NHM

		case matchFamilyModel(info, 0x06, 0x25),
			matchFamilyModel(info, 0x06, 0x2C),
			matchFamilyModel(info, 0x06, 0x2F):
			return INTEL_WSM

		case matchFamilyModel(info, 0x06, 0x2A),
			matchFamilyModel(info, 0x06, 0x2D):
			return INTEL_SNB

		case matchFamilyModel(info, 0x06, 0x3A),
			matchFamilyModel(info, 0x06, 0x3E):
			return INTEL_IVB

		case matchFamilyModel(info, 0x06, 0x3C),
			matchFamilyModel(info, 0x06, 0x3F),
			matchFamilyModel(info, 0x06, 0x45),
			matchFamilyModel(info, 0x06, 0x46):
			return INTEL_HSW

		case matchFamilyModel(info, 0x06, 0x3D),
			matchFamilyModel(info, 0x06, 0x47),
			matchFamilyModel(info, 0x06, 0x4F),
			matchFamilyModel(info, 0x06, 0x56):
			return INTEL_BDW

		case matchFamilyModel(info, 0x06, 0x4E),
			matchFamilyModel(info, 0x06, 0x55),
			matchFamilyModel(info, 0x06, 0x5E):
			return INTEL_SKL

		case matchFamilyModel(info, 0x06, 0x66):
			return INTEL_CNL

		case matchFamilyModel(info, 0x06, 0x7D), // client
			matchFamilyModel(info, 0x06, 0x7E), // client
			matchFamilyModel(info, 0x06, 0x9D), // NNP-I
			matchFamilyModel(info, 0x06, 0x6A), // server
			matchFamilyModel(info, 0x06, 0x6C): // server
			return INTEL_ICL

		case matchFamilyModel(info, 0x06, 0x8C),
			matchFamilyModel(info, 0x06, 0x8D):
			return INTEL_TGL

		case matchFamilyModel(info, 0x06, 0x8F):
			return INTEL_SPR

		case matchFamilyModel(info, 0x06, 0x8E):
			switch info.Stepping {
			case 9:
				return INTEL_KBL
			case 10:
				return INTEL_CFL
			case 11:
				return INTEL_WHL
			case 12:
				return INTEL_CML
			default:
				return X86_UNKNOWN
			}

		case matchFamilyModel(info, 0x06, 0x9E):
			if info.Stepping > 9 {
				return INTEL_CFL
			}
			return INTEL_KBL

		case matchFamilyModel(info, 0x06, 0x97),
			matchFamilyModel(info, 0x06, 0x9A):
			return INTEL_ADL

		case matchFamilyModel(info, 0x06, 0xA5):
			return INTEL_CML

		case matchFamilyModel(info, 0x06, 0xA7):
			return INTEL_RCL

		case matchFamilyModel(info, 0x06, 0x85):
			return INTEL_KNIGHTS_M

		case matchFamilyModel(info, 0x06, 0x57):
			return INTEL_KNIGHTS_L

		case matchFamilyModel(info, 0x0B, 0x00):
			return INTEL_KNIGHTS_F

		case matchFamilyModel(info, 0x0B, 0x01):
			return INTEL_KNIGHTS_C

		case matchFamilyModel(info, 0x0F, 0x01),
			matchFamilyModel(info, 0x0F, 0x02),
			matchFamilyModel(info, 0x0F, 0x03),
			matchFamilyModel(info, 0x0F, 0x04),
			matchFamilyModel(info, 0x0F, 0x06):
			return INTEL_NETBURST
		}

	case cpuid.AMD:
		switch {
		case matchFamilyModel(info, 0xF, 0x04),
			matchFamilyModel(info, 0xF, 0x05),
			matchFamilyModel(info, 0xF, 0x07),
			matchFamilyModel(info, 0xF, 0x08),
			matchFamilyModel(info, 0xF, 0x0C),
			matchFamilyModel(info, 0xF, 0x0E),
			matchFamilyModel(info, 0xF, 0x0F),
			matchFamilyModel(info, 0xF, 0x14),
			matchFamilyModel(info, 0xF, 0x15),
			matchFamilyModel(info, 0xF, 0x17),
			matchFamilyModel(info, 0xF, 0x18),
			matchFamilyModel(info, 0xF, 0x1B),
			matchFamilyModel(info, 0xF, 0x1C),
			matchFamilyModel(info, 0xF, 0x1F),
			matchFamilyModel(info, 0xF, 0x21),
			matchFamilyModel(info, 0xF, 0x23),
			matchFamilyModel(info, 0xF, 0x24),
			matchFamilyModel(info, 0xF, 0x25),
			matchFamilyModel(info, 0xF, 0x27),
			matchFamilyModel(info, 0xF, 0x2B),
			matchFamilyModel(info, 0xF, 0x2C),
			matchFamilyModel(info, 0xF, 0x2F),
			matchFamilyModel(info, 0xF, 0x41),
			matchFamilyModel(info, 0xF, 0x43),
			matchFamilyModel(info, 0xF, 0x48),
			matchFamilyModel(info, 0xF, 0x4B),
			matchFamilyModel(info, 0xF, 0x4C),
			matchFamilyModel(info, 0xF, 0x4F),
			matchFamilyModel(info, 0xF, 0x5D),
			matchFamilyModel(info, 0xF, 0x5F),
			matchFamilyModel(info, 0xF, 0x68),
			matchFamilyModel(info, 0xF, 0x6B),
			matchFamilyModel(info, 0xF, 0x6F),
			matchFamilyModel(info, 0xF, 0x7F),
			matchFamilyModel(info, 0xF, 0xC1):
			return AMD_HAMMER

		case matchFamilyModel(info, 0x10, 0x02),
			matchFamilyModel(info, 0x10, 0x04),
			matchFamilyModel(info, 0x10, 0x05),
			matchFamilyModel(info, 0x10, 0x06),
			matchFamilyModel(info, 0x10, 0x08),
			matchFamilyModel(info, 0x10, 0x09),
			matchFamilyModel(info, 0x10, 0x0A):
			return AMD_K10

		case matchFamilyModel(info, 0x11, 0x03):
			return AMD_K11

		case matchFamilyModel(info, 0x12, 0x01):
			return AMD_K12

		case matchFamilyModel(info, 0x14, 0x00),
			matchFamilyModel(info, 0x14, 0x01),
			matchFamilyModel(info, 0x14, 0x02):
			return AMD_BOBCAT

		case matchFamilyModel(info, 0x15, 0x01):
			return AMD_BULLDOZER

		case matchFamilyModel(info, 0x15, 0x02),
			matchFamilyModel(info, 0x15, 0x11),
			matchFamilyModel(info, 0x15, 0x13):
			return AMD_PILEDRIVER

		case matchFamilyModel(info, 0x15, 0x30),
			matchFamilyModel(info, 0x15, 0x38):
			return AMD_STREAMROLLER

		case matchFamilyModel(info, 0x15, 0x60),
			matchFamilyModel(info, 0x15, 0x65),
			matchFamilyModel(info, 0x15, 0x70):
			return AMD_EXCAVATOR

		case matchFamilyModel(info, 0x16, 0x00):
			return AMD_JAGUAR

		case matchFamilyModel(info, 0x16, 0x30):
			return AMD_PUMA

		case matchFamilyModel(info, 0x17, 0x01),
			matchFamilyModel(info, 0x17, 0x11),
			matchFamilyModel(info, 0x17, 0x18),
			matchFamilyModel(info, 0x17, 0x20):
			return AMD_ZEN

		case matchFamilyModel(info, 0x17, 0x08):
			return AMD_ZEN_PLUS

		case matchFamilyModel(info, 0x17, 0x31),
			matchFamilyModel(info, 0x17, 0x47),
			matchFamilyModel(info, 0x17, 0x60),
			matchFamilyModel(info, 0x17, 0x68),
			matchFamilyModel(info, 0x17, 0x71),
			matchFamilyModel(info, 0x17, 0x90),
			matchFamilyModel(info, 0x17, 0x98):
			return AMD_ZEN2

		case matchFamilyModel(info, 0x19, 0x00),
			matchFamilyModel(info, 0x19, 0x01),
			matchFamilyModel(info, 0x19, 0x08),
			matchFamilyModel(info, 0x19, 0x21),
			matchFamilyModel(info, 0x19, 0x30),
			matchFamilyModel(info, 0x19, 0x40),
			matchFamilyModel(info, 0x19, 0x44),
			matchFamilyModel(info, 0x19, 0x50):
			return AMD_ZEN3

		case matchFamilyModel(info, 0x19, 0x10):
			return AMD_ZEN4
		}
	}

	return X86_UNKNOWN
}

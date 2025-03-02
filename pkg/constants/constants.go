package constants

import (
	"os"
)

const DefaultVPCCIDR = "10.0.0.0/16"
const DefaultCIDRPrefix = 24

var HomeDir, _ = os.UserHomeDir()
var QEFLAG string = os.Getenv("QE_FLAG")

const (
	HTTPConflict                   = 409
	AWSCredentialsFileRelativePath = "/.aws/credentials"
	AdminPolicyArn                 = "arn:aws:iam::aws:policy/AdministratorAccess"
	DefaultAWSCredentialUser       = "default"
	DefaultAWSRegion               = "us-east-2"
	DefaultAWSZone                 = "a"
)

var AWSCredentialsFilePath = HomeDir + AWSCredentialsFileRelativePath

const (
	StageENVTrustedRole    = "arn:aws:iam::644306948063:role/RH-Managed-OpenShift-Installer"
	StageIssuerTrustedRole = "arn:aws:sts::644306948063:assumed-role/RH-Managed-OpenShift-Installer/OCM"
	ProdENVTrustedRole     = "arn:aws:iam::710019948333:role/RH-Managed-OpenShift-Installer"
)

const (
	STSRoleName         = "Installer"
	STSSupportRoleName  = "Support"
	STSMasterRoleName   = "ControlPlane"
	STSWorkerRoleName   = "Worker"
	STSClusterRegion    = "us-west-1"
	DefaultAWSAccountID = "301721915996"
)

const (
	RouteDestinationCidrBlock = "0.0.0.0/0"
	SubnetNamePrefix          = "sdqe-sdk-"
	VpcName                   = "sdqe-sdk-vpc"

	CreationPrivateSelector = "private"
	CreationPublicSelector  = "public"
	CreationPairSelector    = "pair"
	CreationMultiSelector   = "multi"

	VpcDnsHostnamesAttribute = "DnsHostnames"
	VpcDnsSupportAttribute   = "DnsSupport"
	NetworkResourceFileName  = "resource.json"

	TCPProtocol = "tcp"
	UDPProtocol = "udp"

	SecurityGroupName        = "odf-sec-group"
	SecurityGroupDescription = "security group for ms"

	ProxySecurityGroupName        = "proxy-sg"
	AdditionalSecurityGroupName   = "ocm-additional-sg"
	ProxySecurityGroupDescription = "security group for proxy"

	QEFlagKey = "ocm_qe_flag"

	// Proxy related
	ProxyName       = "ocm-proxy"
	InstanceKeyName = "openshift-qe"
	AWSInstanceUser = "ec2-user"
	BastionName     = "ocm-bastion"
)

var ProxyImageMap = map[string]string{
	"us-west-2":      "ami-03b82d95dbe67072d",
	"ap-northeast-1": "ami-0517f6ca1da98f337",
}
var BastionImageMap = map[string]string{
	"us-east-1":      "ami-01c647eace872fc02",
	"us-east-2":      "ami-00a9282ce3b5ddfb1",
	"us-west-1":      "ami-0f1ee917b10382dea",
	"ap-southeast-1": "ami-0db1894e055420bc0",
	"us-west-2":      "ami-0b2b4f610e654d9ac",
	"ap-northeast-1": "ami-0a21e01face015dd9",
}

var DefaultSGIngressPorts = []map[string]int32{
	{"fromPort": 6789, "toPort": 6789},
	{"fromPort": 6800, "toPort": 7300},
	{"fromPort": 3300, "toPort": 3300},
	{"fromPort": 9283, "toPort": 9283},
	{"fromPort": 31659, "toPort": 31659},
}

const (
	PublicSubNetTagKey   = "PublicSubnet"
	PublicSubNetTagValue = "true"
)

const (
	PrivateLBTag = "kubernetes.io/role/internal-elb"
	PublicLBTag  = "kubernetes.io/role/elb"
	// valid LB Tag is empty or 1
	LBTagValue = ""
)
